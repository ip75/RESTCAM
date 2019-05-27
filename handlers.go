package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Picture is a picture
type Picture struct {
	Datetime string `json:"datetime"`
	Image    []byte `json:"image"`
}

func (context *Context) getPicture(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	offBuffer := make([]byte, 100)
	context.sharedFile.Seek(0, 0)
	context.sharedFile.Read(offBuffer)

	w.Write([]byte("send picture with time " + vars["time"] + " now: " + time.Now().String() + "\n"))

	pic := Picture{Datetime: time.Now().String(), Image: offBuffer}
	if err := json.NewEncoder(w).Encode(pic); err != nil {
		panic(err)
	}
}

func getCircleBufferInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
