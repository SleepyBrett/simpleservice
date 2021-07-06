package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	identifier := os.Getenv("HOSTNAME")
	version := os.Getenv("VERSION")
	//sleeptime := os.Getenv("SLEEPTIME")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintf(w, "pod: %s version: %s", identifier, version)
	})


	http.ListenAndServe(":8080", nil)
}