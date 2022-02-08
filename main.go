package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/gorilla/handlers"
)

type Message struct {
	Message    string
	StatusCode int
}

func init() {
	xray.Configure(xray.Config{
		DaemonAddr:     "127.0.0.1:2000", // default
		ServiceVersion: "1.2.3",
	})
}

func handler(w http.ResponseWriter, r *http.Request) {

	json_message := os.Getenv("JSON_MESSAGE")

	message := Message{json_message, 200}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	fmt.Fprintf(os.Stdout, "Starting listening on Port 8080\n")
	http.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler)))
	http.HandleFunc("/health", healthcheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
