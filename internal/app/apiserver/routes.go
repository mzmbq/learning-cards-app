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

		r.Route("/user", func(r chi.Router) {
			r.Get("/whoami", MakeHandler(s.handleUserWhoami()))
			r.Get("/sign-out", MakeHandler(s.handlerUserSignOut()))
		})

		r.Route("/deck", func(r chi.Router) {
			r.Post("/create", MakeHandler(s.handleDeckCreate()))
			r.Delete("/delete/{id}", MakeHandler(s.handleDeckDelete()))
			r.Post("/rename/{id}", MakeHandler(s.handleDeckRename()))
			r.Get("/list-cards/{id}", MakeHandler(s.handleDeckListCards()))
			r.Get("/list", MakeHandler(s.handleDecksList()))
		})

		r.Route("/card", func(r chi.Router) {
			r.Post("/create", MakeHandler(s.handleCardCreate()))
			r.Post("/update/{id}", MakeHandler(s.handleCardUpdate()))
			r.Delete("/delete/{id}", MakeHandler(s.handleCardDelete()))
		})

		r.Route("/study", func(r chi.Router) {
			r.Get("/get-card/{deck_id}", MakeHandler(s.handleStudyGetCard()))
			r.Post("/submit/{card_id}", MakeHandler(s.handleStudySubmit()))
		})

		r.Get("/define/{dict}/{word}", MakeHandler(s.handleDefine()))
		r.Get("/search/{dict}/{word}", MakeHandler(s.handleSearch()))
	})

	s.mux.Mount("/api", apiRouter)
}
