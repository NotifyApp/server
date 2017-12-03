package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var upgrader = websocket.Upgrader{}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", header)
	mux.HandleFunc("/", http.NotFound)
	s := &http.Server{
		Addr:           ":3000",
		Handler:        cors.Default().Handler(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func header(w http.ResponseWriter, r *http.Request) {
	jsonResp := struct {
		Version string `json:"version"`
		Info    string `json:"info"`
	}{
		"1.0.0",
		"Welcome on NotifyApp server",
	}

	resp, err := json.Marshal(jsonResp)
	if err != nil {
		log.Fatal("Error marshal json")
	}
	w.Header().Set("Server", "Notify")
	w.Header().Set("Content-type", "text/json")
	w.WriteHeader(200)
	w.Write(resp)
}
