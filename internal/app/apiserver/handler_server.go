package apiserver

import "net/http"

func (s *server) handleHealthcheck() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
}
