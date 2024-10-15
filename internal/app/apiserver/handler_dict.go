package apiserver

import (
	"errors"
	"fmt"
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
func (s *server) handleSearch() APIFunc {
	type response struct {
		Suggestions []string `json:"suggestions"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		dictName := r.PathValue("dict")
		word := r.PathValue("word")

		word, err := validateWord(word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		d, err := dict.New(dictName)
		if err != nil {
			return InvalidRequestData(map[string]string{"dict": fmt.Sprintf("dictionaray %v not found", dictName)})
		}

		suggs, err := d.Search(word)
		if err != nil {
			if err == dict.ErrTooManyRequests {
				return NewAPIError(http.StatusTooManyRequests, "too many requests")
			}
			return err
		}

		resp := &response{
			Suggestions: suggs,
		}
		return WriteJSON(w, http.StatusOK, resp)
	}

}

// Get the definition of a word
func (s *server) handleDefine() APIFunc {
	type response struct {
		Definitions []dict.Entry `json:"definitions"`
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		dictName := r.PathValue("dict")
		word := r.PathValue("word")
		// TODO: validate dictName and word

		d, err := dict.New(dictName)
		if err != nil {
			return InvalidRequestData(map[string]string{"dict": err.Error()})
		}

		defs, err := d.Define(word)
		if err != nil {
			if err == dict.ErrTooManyRequests {
				return NewAPIError(http.StatusTooManyRequests, err.Error())
			}
			return err
		}

		resp := &response{
			Definitions: defs,
		}
		return WriteJSON(w, http.StatusOK, resp)
	}

}
