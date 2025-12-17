package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LuisanaMTDev/spaced_learning/server/controllers"
	"github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"
	"github.com/LuisanaMTDev/spaced_learning/server/frontend/views"
	"github.com/LuisanaMTDev/spaced_learning/server/helpers"
	"github.com/LuisanaMTDev/spaced_learning/server/middlewares"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {

	godotenv.Load()
	platform := os.Getenv("PLATFORM")
	var db *sql.DB
	var err error

	if platform == "PROD" {
		dbURL := os.Getenv("DB_URL_PROD")
		db, err = sql.Open("libsql", dbURL)
		if err != nil {
			log.Printf("ERROR: while opening db: %v", err)
			return
		}
	} else {
		dbURL := os.Getenv("DB_URL_DEV")
		db, err = sql.Open("sqlite", dbURL)
		if err != nil {
			log.Printf("ERROR: while opening db: %v", err)
			return
		}
	}

	dbQueries := gosql_queries.New(db)
	handler := http.NewServeMux()
	serverConfig := helpers.ServerConfig{DBQueries: dbQueries, Platform: platform}
	server := http.Server{Handler: handler, Addr: ":8090"}
	log.Printf("Running platfotm: %s", serverConfig.Platform)

	//End points
	handler.Handle("GET /app/", http.StripPrefix("/app/", middlewares.ExcludeFiles(
		http.FileServer(http.Dir("./frontend/assets/")), []string{"css/input.css", "images/favicon.svg"},
	)))

	handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		client := r.Header.Get("SL-Client-Type")

		if strings.Contains(client, "SL-CLI") {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("This endpoint isn't intended to work with the CLI."))
		}

		err := views.Index().Render(r.Context(), w)
		if err != nil {
			log.Printf("ERROR: while sending main page: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	handler.HandleFunc("POST /lesson/add", controllers.AddLesson(&serverConfig))

	log.Fatal(server.ListenAndServe())
}
