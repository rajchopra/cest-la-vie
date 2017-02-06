package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DecodeString : Function to decode the base64encoded string to normal string
func DecodeString(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	rawString := params["str"]
	sDec, _ := base64.StdEncoding.DecodeString(rawString)
	fmt.Fprintf(w, "Hi there, Result is %s !!", sDec)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/decode/{str}", DecodeString).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", router))
}