package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uuidWithHyphen := uuid.NewRandom()
		uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
		fmt.Fprintf(w, "Welcome to my website!\n")
		fmt.Fprintf(w, uuid)
	})

	fmt.Print("Starting server in port 8080...\n")

	http.ListenAndServe(":8080", nil)
}
