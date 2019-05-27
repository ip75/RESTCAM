package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Context is context
type Context struct {
	sharedFile *os.File
}

func main() {

	mmfNamePtr := flag.String("mmf", "camera.flow", "memory mapped file to exchange pictures")

	sharedFile, err := os.OpenFile(*mmfNamePtr, os.O_RDONLY, 0600)
	context := &Context{sharedFile: sharedFile}
	defer sharedFile.Close()

	router := mux.NewRouter()
	router.HandleFunc("/pictures", getCircleBufferInfo).Methods("GET")
	router.HandleFunc("/picture/{time:[0-9]{1,2}:[0-9]{2}:[0-9]{2}", context.getPicture).Methods("GET")

	if err != nil {
		log.Fatal(err)
	}
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
