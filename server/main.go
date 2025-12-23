package main

import (
	"database/sql"
	"net/http"
	"os"
	"strings"

	"github.com/LuisanaMTDev/spaced_learning/server/controllers"
	"github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"
	"github.com/LuisanaMTDev/spaced_learning/server/frontend/views"
	"github.com/joho/godotenv"
	// "github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

func main() {

	godotenv.Load()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	platform := os.Getenv("PLATFORM")
	var db *sql.DB
	var err error

	if platform == "PROD" {
		dbURL := os.Getenv("DB_URL_PROD")
		db, err = sql.Open("libsql", dbURL)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("ERROR: while opening db.")
			return
		}
	} else {
		dbURL := os.Getenv("DB_URL_DEV")
		db, err = sql.Open("sqlite", dbURL)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("ERROR: while opening db.")
			return
		}
	}

	dbQueries := gosql_queries.New(db)
	handler := http.NewServeMux()
	serverConfig := controllers.NewServerConfig(dbQueries, platform)

	//End points
	handler.Handle("GET /app/", http.StripPrefix("/app/", http.FileServer(http.Dir("./frontend/assets/"))))

	handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		client := r.Header.Get("SL-Client-Type")

		if strings.Contains(client, "SL-CLI") {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("This endpoint isn't intended to work with the CLI."))
		}

		err := views.Index().Render(r.Context(), w)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("ERROR: while sending main page")
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	handler.HandleFunc("POST /lesson/add", serverConfig.AddLesson)

	server := http.Server{Handler: handler, Addr: ":8090"}
	log.Info().Str("running_platfotm", serverConfig.Platform).Msg("Running...")
	err = server.ListenAndServe()
	log.Fatal().AnErr("server_error", err)
}
