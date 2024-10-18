package main

import (
	"fmt"
	"log"

	"github.com/hash-f/go-apis/game"
)

func main() {
	fmt.Println("Starting...")
	mux := newMux()
	gs := game.NewGameStore()
	gs.AddFakeGame()
	api := newApi(":3000", mux, &gs)
	api.registerHandlers()
	log.Fatal(api.server.ListenAndServe())
}
