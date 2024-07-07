package dict

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	searchURL = "https://en.wiktionary.org/w/index.php?fulltext=Search&search="
	defineURL = "https://en.wiktionary.org/api/rest_v1/page/definition/"
)

type Wiktionary struct{}

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

func (w *Wiktionary) Search(word string) ([]string, error) {
	results := make([]string, 0)

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

	doc.Find(".mw-search-result-heading").Each(func(i int, s *goquery.Selection) {
		result := strings.TrimSpace(s.Text())
		if result != "" {
			results = append(results, result)
		}
	})

	return results, nil
}

func (w *Wiktionary) Define(word string) ([]Entry, error) {
	ets := make([]Entry, 0)

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
