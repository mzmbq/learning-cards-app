package apiserver

import (
	"context"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/sqlstore"
	"golang.org/x/time/rate"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL, time.Duration(config.DatabaseTimeout))
	if err != nil {
		return err
	}
	store := sqlstore.New(db)

	sessionKey, err := hex.DecodeString(config.SessionKey)
	log.Println("Session key length:", len(sessionKey))
	if err != nil {
		return err
	}
	sessionStore := sessions.NewCookieStore(sessionKey)

	var globalLimiter *rate.Limiter = nil
	if config.GlobaRatelLimit.Enabled {
		globalLimiter = rate.NewLimiter(rate.Limit(config.GlobaRatelLimit.Rps), config.GlobaRatelLimit.Burst)
	}

	var userLimiter *rate.Limiter = nil
	if config.GlobaRatelLimit.Enabled {
		log.Println("IP-based Rate Limiting is enabled")
		userLimiter = rate.NewLimiter(rate.Limit(config.UserRateLimit.Rps), config.UserRateLimit.Burst)
	}

	srv := newServer(store, sessionStore, config.CORSOrigins, globalLimiter, userLimiter)

	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(dbURL string, timeout time.Duration) (*sql.DB, error) {
	log.Println("Connecting to database")
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		log.Println("Failed to connect to database")
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
