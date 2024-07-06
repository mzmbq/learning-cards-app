package dictionary

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	searchURL = "https://en.wiktionary.org/w/index.php?fulltext=Search&search="
	defineURL = "https://en.wiktionary.org/wiki/"
)

type Wiktionary struct {
	// lang        string
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

	// Parse
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".use-with-mention").Each(func(_ int, s *goquery.Selection) {
		et := Entry{
			Word:       word,
			Definition: s.Text(),
		}

		s.Find(".h-usage-example").Each(func(_ int, sel *goquery.Selection) {
			et.Examples = append(et.Examples, sel.Text())
		})

		ets = append(ets, et)
	})

	return ets, nil
}
