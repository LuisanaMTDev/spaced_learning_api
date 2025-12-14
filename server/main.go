package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"
	"github.com/LuisanaMTDev/spaced_learning/server/helpers"
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
			log.Printf("Error while opening db: %v", err)
			return
		}
	} else {
		dbURL := os.Getenv("DB_URL_DEV")
		db, err = sql.Open("sqlite", dbURL)
		if err != nil {
			log.Printf("Error while opening db: %v", err)
			return
		}
	}

	dbQueries := gosql_queries.New(db)
	handler := http.NewServeMux()
	serverConfig := helpers.ServerConfig{DBQueries: dbQueries, Platform: platform}
	server := http.Server{Handler: handler, Addr: ":8090"}
	log.Printf("Running platfotm: %s", serverConfig.Platform)

	//End points
	handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World."))
	})

	log.Fatal(server.ListenAndServe())
}
