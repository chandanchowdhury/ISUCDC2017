package main

import (
	"fmt"
	"github.com/cipherboy/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	var address string = ":8000"
	var err error

	if len(os.Args) > 2 {
		fmt.Println("Usage:", os.Args[0], "[address:port]")
	}

	if len(os.Args) >= 2 {
		address = os.Args[1]
	}

	var api_version string = "/api"
	fs := http.FileServer(http.Dir("../www"))

	routes := mux.NewRouter()
	routes.PathPrefix(api_version).Path("/client").Handler(&ClientHandler{})

	routes.PathPrefix("/").Handler(fs)

	http.Handle("/", routes)

	log.Println("Starting patISEnt www...")
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
