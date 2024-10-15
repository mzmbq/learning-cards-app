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

		r.Get("/decks/list", MakeHandler(s.handleDecksList()))
		r.Post("/deck/create", MakeHandler(s.handleDeckCreate()))
		r.Get("/deck/delete/{id}", MakeHandler(s.handleDeckDelete()))
		r.Get("/deck/list-cards/{id}", MakeHandler(s.handleDeckListCards()))
		r.Post("/deck/rename/{id}", MakeHandler(s.handleDeckRename()))

		r.Post("/card/create", MakeHandler(s.handleCardCreate()))
		r.Post("/card/update/{id}", MakeHandler(s.handleCardUpdate()))
		r.Get("/card/delete/{id}", MakeHandler(s.handleCardDelete()))

		r.Get("/study/get-card/{deck_id}", MakeHandler(s.handleStudyGetCard()))
		r.Post("/study/submit/{card_id}", MakeHandler(s.handleStudySubmit()))

		r.Get("/define/{dict}/{word}", MakeHandler(s.handleDefine()))
		r.Get("/search/{dict}/{word}", MakeHandler(s.handleSearch()))
	})

	s.mux.Mount("/api", apiRouter)
}
