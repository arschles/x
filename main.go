package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const DatastoreKey = "datastore"
const WidthKey = "width"
const HeightKey = "height"
const FormatKey = "Format"

type Datastore string
type Width string
type Height string
type Format string

func getQSVals(req *http.Request) (Datastore, Width, Height, Format) {
	vals := req.URL.Query()
	ds := Datastore(vals.Get(DatastoreKey))
	w := Width(vals.Get(WidthKey))
	h := Height(vals.Get(HeightKey))
	f := Format(vals.Get(FormatKey))
	return ds, w, h, f
}

func main() {
	port := flag.Int("port", 8080, "the port to run on")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/v1/thumb", thumbHandler).
		Methods("POST").
		Queries(DatastoreKey, "").
		Queries(WidthKey, "").
		Queries(HeightKey, "").
		Queries(FormatKey, "")
	router.HandleFunc("/v1/gallery", galleryHandler).
		Methods("POST").
		Queries(DatastoreKey, "").
		Queries(WidthKey, "").
		Queries(HeightKey, "").
		Queries(FormatKey, "")

	http.Handle("/", router)

	hostStr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(hostStr, nil))
}
