package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "v2")
}

func main() {
	log.Println("hello world")
	go func() {
		time.Sleep(time.Second * 30)
		log.Fatal("crash")
	}()

	http.HandleFunc("/", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}
