package main

import (
	"fmt"
	"net/http"

	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/gorilla/mux"
	"github.com/tmax-cloud/image-scanning-webhook/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	port          = 80
	apiPathPrefix = "/webhook"
	clairPrefix   = "/clair"
)

var logWebhook = logf.Log.WithName("webhook")

func main() {
	logf.SetLogger(zap.Logger(true))
	logWebhook.Info("initializing server....")

	router := mux.NewRouter()
	apiRouter := router.PathPrefix(apiPathPrefix).Subrouter()

	//clair/image-scanning
	apiRouter.HandleFunc(clairPrefix, apis.CreateClairLog).Methods("POST")

	//harbor/image-scanning

	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		logWebhook.Error(err, "failed to initialize a server")
	}
}
