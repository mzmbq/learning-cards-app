package sqlstore_test

import (
	"flag"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/apiserver"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/sqlstore"
)

var store *sqlstore.Store

func TestMain(m *testing.M) {
	dbURL := os.Getenv("TESTDB_URL")
	db, err := apiserver.NewDB(dbURL)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	store = sqlstore.New(db)

	flag.Parse()
	exitCode := m.Run()

	defer db.Close()

	os.Exit(exitCode)
}

func TestCardRepository_FindAllByDeckID(t *testing.T) {
	d := model.Deck{
		ID:     1,
		Name:   "Test",
		UserID: 1,
	}

	cards, err := store.Card().FindAllByDeckID(d.ID)
	if err != nil {
		log.Fatal(err)
	}
	if cards == nil {
		log.Fatal("cards in nill")
	}

}
