package apiserver

import (
	"github.com/go-chi/chi/v5"
)

func (s *server) routes() {
	apiRouter := chi.NewRouter()

	apiRouter.Use(s.withLogging, s.withCORS, s.withGlobalRateLimit)

	// Routes that do not require authentication
	apiRouter.Group(func(r chi.Router) {
		r.Get("/health", MakeHandler(s.handleHealthcheck()))
		r.Post("/user/create", MakeHandler(s.handleUserCreate()))
		r.Post("/user/auth", MakeHandler(s.handleUserAuth()))
	})

	// Routes that require authentication
	apiRouter.Group(func(r chi.Router) {
		r.Use(s.withAuth, s.withUserRateLimit)

		r.Get("/user/whoami", MakeHandler(s.handleUserWhoami()))
		r.Get("/user/signout", MakeHandler(s.handlerUserSignOut()))

		r.Get("/decks/list", s.handleDecksList())
		r.Post("/deck/create", s.handleDeckCreate())
		r.Get("/deck/delete/{id}", s.handleDeckDelete())
		r.Get("/deck/list-cards/{id}", s.handleDeckListCards())
		r.Post("/deck/rename/{id}", s.handleDeckRename())

		r.Post("/card/create", s.handleCardCreate())
		r.Post("/card/update/{id}", s.handleCardUpdate())
		r.Get("/card/delete/{id}", s.handleCardDelete())

		r.Get("/study/get-card/{deck_id}", s.handleStudyGetCard())
		r.Post("/study/submit/{card_id}", s.handleStudySubmit())

		r.Get("/define/{dict}/{word}", s.handleDefine())
		r.Get("/search/{dict}/{word}", s.handleSearch())
	})

	s.mux.Mount("/api", apiRouter)
}
