package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jitaeyun/image-scanning-webhook/pkg/apis"
)

const (
	port          = 80
	apiPathPrefix = "/webhook"
	clairPrefix   = "/clair"
)

func main() {
	log.Println("initializing server....")

	router := mux.NewRouter()
	apiRouter := router.PathPrefix(apiPathPrefix).Subrouter()

	//clair/image-scanning
	apiRouter.HandleFunc(clairPrefix, apis.CreateClairLog).Methods("POST")

	//harbor/image-scanning

	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err, "failed to initialize a server")
	}
}
