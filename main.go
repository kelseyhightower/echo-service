package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	log.Println("Starting echo service ...")

	httpListenPort := os.Getenv("PORT")
	if httpListenPort == "" {
		httpListenPort = "8080"
	}

	hostPort := net.JoinHostPort("0.0.0.0", httpListenPort)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		return
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "", 500)
			return
		}
		defer r.Body.Close()

		var message Message
		err = json.Unmarshal(data, &message)
		if err != nil {
			log.Println(err)
			http.Error(w, "", 500)
			return
		}

		responseData, err := json.MarshalIndent(message, "", "  ")
		if err != nil {
			log.Println(err)
			http.Error(w, "", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(responseData))
	})

	s := &http.Server{
		Addr:    hostPort,
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
