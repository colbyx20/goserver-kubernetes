package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	id int
}

func main() {
	DataStore := Data{
		id: 0,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Docker!!!")
	})

	mux.HandleFunc("GET /task/{id}", func(w http.ResponseWriter, r *http.Request) {
		idstr := r.PathValue("id")
		id, err := strconv.Atoi(r.PathValue("id"))

		DataStore.id = id

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hanlding task with id=%v\n", idstr)
	})

	fmt.Println("Listening....")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
