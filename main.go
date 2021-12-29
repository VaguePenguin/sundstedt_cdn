package main

import (
	// std libs
	"log"
	"net/http"
	"os"
	"time"

	// external libs
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const service string = "shared"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.PathPrefix("/js/").Handle(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	r.PathPrefix("/css/").Handle(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.PathPrefix("/img/").Handle(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	c := handlers.AllowedOrigins([]string{ "https://evenson.sundstedt.us", "https://iam.sundstedt.us" }

	port := os.Genenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr: "0.0.0.0:" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Second * 60,
		Handler: handlers.CORS(c)(r),
	}

	log.Fatal(srv.ListenAndServe())
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello CDN")
}


