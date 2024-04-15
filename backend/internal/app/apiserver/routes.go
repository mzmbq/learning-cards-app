package apiserver

func (s *server) routes() {
	s.router.HandleFunc("GET /", handleRoot)

	s.router.Handle("POST /api/user/create", withLogging(withCORS(s.handleUserCreate())))
	s.router.Handle("GET /api/user/{id}", withLogging(withCORS((s.handleUserFind()))))

	s.router.Handle("GET /api/word", s.handleWordDefine())

	s.router.Handle("POST /api/deck", s.handleDeckCreate())
	s.router.Handle("GET /api/deck/{id}", s.handleDeckGet())

	s.router.Handle("POST /api/card", s.handleCardCreate())
	s.router.Handle("GET /api/deck/{deckId}/card", s.handleCardLearn())
	s.router.Handle("POST /api/card/{id}", s.handleCardUpdate())

}
