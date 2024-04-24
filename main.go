package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const INDEX = `<!DOCTYPE html>
<html lang="en">
<head>
    <title>Home Page</title>
</head>
<body>
    <h1>Go Sample Code for App Engine</h1>
</body>
</html>`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, INDEX)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
