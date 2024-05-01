
package main

import (
	"log"
	"net/http"
	sw "todoserver/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
