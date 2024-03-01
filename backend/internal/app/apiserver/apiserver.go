package apiserver

import (
	"net/http"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/teststore"
)

func Start(config *Config) error {
	store := teststore.New()

	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}
