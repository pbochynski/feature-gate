package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	knative := os.Getenv("KNATIVE")
	if knative == "" {
		fmt.Fprintf(w, "Hello world!") // send data to client side
	} else {
		fmt.Fprint(w, "Hello Knative world!")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
