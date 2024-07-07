package dict

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/time/rate"
)

const (
	searchURL = "https://en.wiktionary.org/w/index.php?fulltext=Search&search="
	defineURL = "https://en.wiktionary.org/api/rest_v1/page/definition/"
)

var (
	limiter = rate.NewLimiter(200, 200) // 200 req/s
)

type Wiktionary struct{}

func (w *Wiktionary) Search(word string) ([]string, error) {
	if !limiter.Allow() {
		return nil, ErrTooManyRequests
	}

	// Fetch
	resp, err := http.Get(searchURL + word)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	doc.Find(".mw-search-result-heading").Each(func(i int, s *goquery.Selection) {
		result := strings.TrimSpace(s.Text())
		if result != "" {
			results = append(results, result)
		}
	})

	return results, nil
}

type DefinitionsResponse struct {
	Usages []UsageDescription `json:"en"`
}

type UsageDescription struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Description string   `json:"description"`
	Definition  string   `json:"definition"`
	Examples    []string `json:"examples,omitempty"`
}

func (w *Wiktionary) Define(word string) ([]Entry, error) {
	if !limiter.Allow() {
		return nil, ErrTooManyRequests
	}

	// Fetch
	resp, err := http.Get(defineURL + word)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, errors.New("no defintion found")
	}

	res := DefinitionsResponse{}
	log.Println(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	ets := make([]Entry, 0)
	for _, u := range res.Usages {
		for _, d := range u.Definitions {
			et := Entry{
				Word:       word,
				Definition: d.Definition,
				Examples:   d.Examples,
			}
			ets = append(ets, et)
		}
	}

	return ets, nil
}
