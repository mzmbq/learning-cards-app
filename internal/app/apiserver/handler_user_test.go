package apiserver

import (
	"encoding/json"
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
}

func TestServer_handlerUserSignOut(t *testing.T) {
}

func TestServer_handleUserWhoami(t *testing.T) {
}
