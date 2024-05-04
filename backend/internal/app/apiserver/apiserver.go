package apiserver

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	store := sqlstore.New(db)

	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	log.Println("Connecting to database")
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		log.Println("Failed to connect to database")
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
