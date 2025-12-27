package controllers

import (
	"database/sql"
	"os"

	"github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

// This struct will contain the server config, as its name says, and the controllers.
type ServerConfig struct {
	DBQueries   *gosql_queries.Queries
	Platform    string
	OAuthConfig oauth2.Config
}

func NewServerConfig() *ServerConfig {

	platform := os.Getenv("PLATFORM")
	var db *sql.DB
	var err error

	if platform == "PROD" {
		dbURL := os.Getenv("DB_URL_PROD")
		db, err = sql.Open("libsql", dbURL)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("ERROR: while opening db.")
			return nil
		}
	} else {
		dbURL := os.Getenv("DB_URL_DEV")
		db, err = sql.Open("sqlite", dbURL)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("ERROR: while opening db.")
			return nil
		}
	}

	dbQueries := gosql_queries.New(db)
	return &ServerConfig{
		DBQueries: dbQueries,
		Platform:  platform,
		OAuthConfig: oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://github.com/login/oauth/authorize",
				TokenURL: "https://github.com/login/oauth/access_token",
			},
			RedirectURL: "http://localhost:8090/oauth2/callback",
			Scopes:      []string{"user"},
		},
	}
}
