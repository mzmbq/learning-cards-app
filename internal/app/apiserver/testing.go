package apiserver

import (
	"testing"

	"github.com/gorilla/sessions"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/teststore"
)

func TestServer(t *testing.T) *server {
	t.Helper()

	store := teststore.New()
	sessionStore := sessions.NewCookieStore([]byte("test"))
	s := newServer(store, sessionStore, []string{}, nil, nil)
	return s
}
