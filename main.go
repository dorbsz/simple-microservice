package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func indexHandler(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		fmt.Fprintf(w, "Go is alive! with %d", rand.Intn(100))

	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not allowed")
	}

}

func main() {
	fmt.Println("Starting go webserver")
	http.HandleFunc("/go", indexHandler)

	http.ListenAndServe(":80", nil)
}
