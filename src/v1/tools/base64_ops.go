package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// EncodeString : Function to encode the raw string to base64encoded string
func EncodeString(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	rawString := params["str"]
	sEnc := base64.StdEncoding.EncodeToString([]byte(rawString))
	w.Header().Set("Content-Type", "audio/mpeg; charset=utf-8")
	w.Write([]byte("Result : " + sEnc + " !!"))
}

// DecodeString : Function to decode the base64encoded string to normal string
func DecodeString(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	rawString := params["str"]
	sDec, _ := base64.StdEncoding.DecodeString(rawString)
	fmt.Fprintf(w, "Hi there, Result is %s !!", sDec)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/encode/{str}", EncodeString).Methods("GET")
	router.HandleFunc("/decode/{str}", DecodeString).Methods("GET")
	log.Fatal(http.ListenAndServe(":2211", router))
}
