package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/vextor22/go_docker/app/restservice"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	restservice.RegisterHelloWorlds(r)
	restservice.RegisterRedisEndpoints(r)
	restservice.RegisterMongoEndpoints(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App does not have this endpoint")
}
