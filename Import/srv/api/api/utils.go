package main

import (
	"net/http"
)

const isemr_addr string = "http://emr.indianola.isucdc.com:80/"

func headers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Header().Set("Content-type", "application/json")
}

func respondNotImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("{\"status\": \"error\", \"message\": \"Not implemented.\"}"))
}
