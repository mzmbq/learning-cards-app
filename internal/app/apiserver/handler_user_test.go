package apiserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestServer_handleUserCreate(t *testing.T) {
	s := TestServer(t)
	err := s.store.User().Create(&model.User{
		Email: "taken@mail.com",
	})
	assert.Nil(t, err)

	testCases := []struct {
		name         string
		body         map[string]any
		expectedCode int
	}{
		{
			name:         "valid user",
			body:         map[string]any{"email": "valid@test.com", "password": "123456"},
			expectedCode: http.StatusOK,
		},
		{
			name:         "short password",
			body:         map[string]any{"email": "short@test.com", "password": "123"},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "duplicate email",
			body:         map[string]any{"email": "taken@mail.com", "password": "123456"},
			expectedCode: http.StatusConflict,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatalf("Failed to marshal body: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/user/create", strings.NewReader(string(bodyBytes)))
			rr := httptest.NewRecorder()

			s.ServeHTTP(rr, req)
			res := rr.Result()

			assert.NotNil(t, res)
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tc.expectedCode)
		})
	}
}

func TestServer_handleUserAuth(t *testing.T) {
	s := TestServer(t)
	err := s.store.User().Create(&model.User{
		Email:    "alice@mail.com",
		Password: "secret123",
	})
	assert.Nil(t, err)

	testCases := []struct {
		name         string
		body         map[string]any
		expectedCode int
	}{
		{
			name:         "wrong password",
			body:         map[string]any{"email": "alice@mail.com", "password": "nicetry123"},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "correct password",
			body:         map[string]any{"email": "alice@mail.com", "password": "secret123"},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not existing user",
			body:         map[string]any{"email": "bob@mail.com", "password": "secret123"},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes, err := json.Marshal(tc.body)
			assert.Nil(t, err, "Failed to marshal body: %v", err)

			req := httptest.NewRequest(http.MethodPost, "/api/user/auth", strings.NewReader(string(bodyBytes)))
			rr := httptest.NewRecorder()

			s.ServeHTTP(rr, req)
			res := rr.Result()

			assert.NotNil(t, res)
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tc.expectedCode)
		})
	}
}

func TestServer_handlerUserSignOut(t *testing.T) {
	s := TestServer(t)
	u := &model.User{
		Email:    "alice@mail.com",
		Password: "secret123",
	}
	err := s.store.User().Create(u)
	assert.Nil(t, err)

	// Login
	body := map[string]string{
		"email":    "alice@mail.com",
		"password": "secret123",
	}
	bodyBytes, err := json.Marshal(body)
	assert.Nil(t, err, "Failed to marshal body: %v", err)

	req := httptest.NewRequest(http.MethodPost, "/api/user/auth", strings.NewReader(string(bodyBytes)))
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	res := rr.Result()
	cookies := res.Cookies()
	assert.NotNil(t, res)
	res.Body.Close()

	// Call whoami
	req = httptest.NewRequest(http.MethodGet, "/api/user/whoami", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rr = httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	res = rr.Result()
	assert.NotNil(t, res)
	res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	// Logout
	req = httptest.NewRequest(http.MethodGet, "/api/user/signout", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rr = httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	res = rr.Result()
	newCookies := res.Cookies()
	assert.NotNil(t, res)
	res.Body.Close()

	// Call whoami again, now with new cookies
	req = httptest.NewRequest(http.MethodGet, "/api/user/whoami", nil)
	for _, c := range newCookies {
		req.AddCookie(c)
	}
	rr = httptest.NewRecorder()

	s.ServeHTTP(rr, req)
	res = rr.Result()
	assert.NotNil(t, res)
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusUnauthorized)

}

func TestServer_handleUserWhoami(t *testing.T) {
	s := TestServer(t)
	u := &model.User{
		Email:    "alice@mail.com",
		Password: "secret123",
	}
	err := s.store.User().Create(u)
	assert.Nil(t, err)

	testCases := []struct {
		name         string
		userID       any
		expectedCode int
	}{
		{
			name:         "valid user",
			userID:       u.ID,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid format",
			userID:       "a1",
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "not existing user",
			userID:       123,
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), ctxKeyUserID, tc.userID)
			log.Println("user id", tc.userID)

			req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			// Call handler directly
			h := MakeHandler(s.handleUserWhoami())
			h(rr, req)
			res := rr.Result()

			log.Println(rr.Body)

			assert.NotNil(t, res)
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tc.expectedCode)
		})
	}
}
