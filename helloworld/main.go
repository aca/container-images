package main

import (
	"fmt"
	"log"
	"net/http"
)

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "v1")
}

func main() {
	http.HandleFunc("/", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
