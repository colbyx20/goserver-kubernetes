package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Docker!!!")
	})

	fmt.Println("Listening....")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
