package apiserver

func (s *server) routes() {
	s.router.HandleFunc("GET /", rootHandler)

	s.router.Handle("POST /api/user/create/", withLogging(withCORS(s.handleUserCreate())))
	s.router.Handle("GET /api/user/{id}", withLogging(withCORS((s.handleUserFind()))))
}
