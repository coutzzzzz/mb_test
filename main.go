package main

import (
	"github.com/coutzzzzz/mb-go-test/internal/daemon"
	"github.com/gorilla/mux"
)

func main() {
	muxServer := mux.NewRouter()
	daemon.ServeHTTP(muxServer)
}
