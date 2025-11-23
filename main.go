package main

import (
	"log"
	"math/rand"
	"net/http"
)

var messages = []string{
	"Now is the time for all good Devops to come the aid of their servers.",
	"Alas poor Altair 8800; I knew it well!",
	"In the beginning there was ARPA Net and its domain was limited.",
	"A blog a day helps keep the hacker away.",
}

var spec = ":8080"

func sendRandomMessage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add(http.CanonicalHeaderKey("Content-Type"), "application/json; charset=utf-8")
	w.Write([]byte(messages[rand.Intn(len(messages))]))
}
func main() {
	http.HandleFunc("/message", sendRandomMessage)
	if err := http.ListenAndServe(spec, nil); err != nil {
		log.Fatalf("Failed to start server: %v on :%v", err, spec)
	}

}
