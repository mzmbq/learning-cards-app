package apiserver

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/dict"
)

var invalidCharsRegex = regexp.MustCompile(`[\/\\:;?!@#$%^&*()\[\]{}<>|~]`)

func validateWord(word string) (string, error) {
	word = strings.TrimSpace(word)

	if invalidCharsRegex.MatchString(word) {
		return "", errors.New("word query contains invalid characters")
	}
	return word, nil
}

// Get suggestions for a word
func (s *server) handleSearch() http.HandlerFunc {
	type response struct {
		Suggestions []string `json:"suggestions"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		dictName := r.PathValue("dict")
		word := r.PathValue("word")

		word, err := validateWord(word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		d, err := dict.New(dictName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		suggs, err := d.Search(word)
		if err != nil {
			if err == dict.ErrTooManyRequests {
				http.Error(w, err.Error(), http.StatusTooManyRequests)
				return
			}
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		resp := &response{
			Suggestions: suggs,
		}
		s.WriteJSON(w, http.StatusOK, resp)
	}

}

// Get the definition of a word
func (s *server) handleDefine() http.HandlerFunc {
	type response struct {
		Definitions []dict.Entry `json:"definitions"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		dictName := r.PathValue("dict")
		word := r.PathValue("word")
		// TODO: validate dictName and word

		d, err := dict.New(dictName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defs, err := d.Define(word)
		if err != nil {
			if err == dict.ErrTooManyRequests {
				http.Error(w, err.Error(), http.StatusTooManyRequests)
				return
			}
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		resp := &response{
			Definitions: defs,
		}
		s.WriteJSON(w, http.StatusOK, resp)
	}

}
