package controllers

import (
	"encoding/json"
	"log"
	"net/http"

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
				log.Printf("Error while decoding json request: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Println(body)
		case "SL-WEB-APP":
			err := r.ParseForm()

			if err != nil {
				log.Printf("Error while parsing request form: %s", err)
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
