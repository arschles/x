package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "flag"
  "log"
  "fmt"
)

func main() {
  port := flag.Int("port", 8080, "the port to run on")
  flag.Parse()

  router := mux.NewRouter()
  router.HandleFunc("/v1/thumb", thumbHandler).
    Methods("POST").
    Queries("datastore", "").
    Queries("width", "").
    Queries("height", "").
    Queries("format", "")
  router.HandleFunc("/v1/gallery", galleryHandler).
    Methods("POST").
    Queries("datastore", "").
    Queries("width", "").
    Queries("height", "").
    Queries("format", "")

  http.Handle("/", router)

  hostStr := fmt.Sprintf(":%d", *port)
  log.Fatal(http.ListenAndServe(hostStr, nil))
}
