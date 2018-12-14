package view

import (
	config "config/server"
	"controller"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type restServer struct {
	controller controller.Controller
	ip         string
	port       string
	router     *mux.Router
}

func NewRestServer(config config.Config, controller controller.Controller) View {
	return &restServer{
		controller: controller,
		ip:         config.Ip(),
		port:       config.Port(),
		router:     mux.NewRouter()}
}

func (server restServer) Run(wg sync.WaitGroup) error {
	log.Println("Starting rest server...")
	server.router.HandleFunc("/games", server.doGetGames).Methods("GET")
	server.router.HandleFunc("/games", server.doNewGame).Methods("POST")
	server.router.HandleFunc("/games/{id}/guesses", server.doNewGuess).Methods("PUT")
	loggedRouter := handlers.LoggingHandler(os.Stdout, server.router)
	http.ListenAndServe(server.ip+":"+server.port, loggedRouter)
	wg.Done()
	return nil
}

func (server restServer) doNewGame(w http.ResponseWriter, r *http.Request) {
	gameInfo, err := server.controller.NewGame()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	gameInfoJson, _ := json.Marshal(gameInfo)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gameInfoJson)
}

func (server restServer) doGetGames(w http.ResponseWriter, r *http.Request) {
	gamesInfo, err := server.controller.GetGames()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	gamesJson, _ := json.Marshal(gamesInfo)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gamesJson)
}

type userGuess struct {
	Guess string
}

func (server restServer) doNewGuess(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var guess userGuess
	err = json.Unmarshal(body, &guess)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	gameInfo, err := server.controller.NewGuess(params["id"], guess.Guess)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	gameInfoJson, _ := json.Marshal(gameInfo)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gameInfoJson)
}
