package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"
	"github.com/LuisanaMTDev/spaced_learning/server/helpers"
)

func AddLesson(sc *helpers.ServerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client := r.Header.Get("SL-Client-Type")

		switch client {
		case "SL-CLI":
			body := AddLessonRequest{}
			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				log.Printf("ERROR: while decoding json request: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Printf("SUCCESS: Decoded request body: %v", body)

			repetitionsDatesAsJson, err := json.Marshal(body.RepetitionsDates)

			if err != nil {
				log.Printf("ERROR: while marshaling repetitions dates on decoded request body: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			repetitionsDatesAsString := string(repetitionsDatesAsJson)

			err = sc.DBQueries.AddLesson(r.Context(), gosql_queries.AddLessonParams{
				Topic:         body.Topic,
				StartedDate:   body.RepetitionsDates[0],
				Json:          repetitionsDatesAsString,
				AmountOfCards: body.AmountOfCards,
			})

			if err != nil {
				log.Printf("ERROR: while adding lesson to the database: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			currentLessons, err := sc.DBQueries.GetAllLessons(r.Context())

			if err != nil {
				log.Printf("ERROR: while getting lessons from the database: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Printf("SUCCESS: Current lessons in the database: %v", currentLessons)

			w.Write([]byte("Request has success."))
		case "SL-WEB-APP":
			err := r.ParseForm()

			if err != nil {
				log.Printf("ERROR: while parsing request form: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid SL-Client-Type header"))
		}
	}
}
