package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hash-f/go-apis/game"
)

type GuessRequest struct {
	Guess string
}

type Api struct {
	mux    *http.ServeMux
	gs     *game.GameStore
	server *http.Server
}

func newApi(addr string, handler *http.ServeMux, gs *game.GameStore) *Api {
	return &Api{
		gs:  gs,
		mux: handler,
		server: &http.Server{
			Addr:           addr,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}

}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func sendMethodNotAllowedResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	response := map[string]bool{"status": false}
	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)

}

func (api *Api) registerNewGameHandler() {
	api.mux.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			sendMethodNotAllowedResponse(w)
			return
		}
		g := game.NewGame()
		(*api.gs)[g.Id] = g
		gameJson, _ := json.Marshal(g)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(gameJson)
	})
}

func (api *Api) registerNewGuessHandler() {
	api.mux.HandleFunc("/game/{id}/guess", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendMethodNotAllowedResponse(w)
			return
		}

		id := r.PathValue("id")
		g, ok := (*api.gs)[id]
		if !ok {
			return
		}

		body := &GuessRequest{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := g.Match(body.Guess)
		responseJson, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	})
}

func (api *Api) registerGetGameHandler() {
	api.mux.HandleFunc("/game/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		g, ok := (*api.gs)[id]
		if !ok {
			return
		}

		responseJson, _ := json.Marshal(g)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	})
}

func (api *Api) registerHandlers() {
	api.registerNewGameHandler()
	api.registerNewGuessHandler()
	api.registerGetGameHandler()
}
