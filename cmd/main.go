package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Session struct {
	Session []string `json:"session"`
}

var (
	session = &Session{}
)

func main() {
	http.HandleFunc("/session", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error in Server: %+v", err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	session.Session = append(session.Session, uuid.NewString())
	str, err := json.Marshal(session)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(str)
}
