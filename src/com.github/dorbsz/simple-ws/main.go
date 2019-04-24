package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		fmt.Fprintf(w, "Go is alive! with %s", request.URL.Hostname())

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
