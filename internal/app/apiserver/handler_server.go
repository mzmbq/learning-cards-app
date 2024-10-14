package apiserver

import "net/http"

func (s *server) handleHealthcheck() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}
