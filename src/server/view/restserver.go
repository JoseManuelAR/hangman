package view

import (
	"config"
	"data"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/controller"

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

const baseUrl string = "/hangman/v1"

func (server restServer) Start(bc chan bool) error {
	log.Println("Starting rest server...")
	server.router.HandleFunc(baseUrl+"/games", server.doGetGames).Methods("GET")
	server.router.HandleFunc(baseUrl+"/games", server.doNewGame).Methods("POST")
	server.router.HandleFunc(baseUrl+"/games/{id}", server.doGetGame).Methods("GET")
	server.router.HandleFunc(baseUrl+"/games/{id}/guesses", server.doNewGuess).Methods("PUT")
	loggedRouter := handlers.LoggingHandler(os.Stdout, server.router)
	http.ListenAndServe(server.ip+":"+server.port, loggedRouter)
	bc <- true
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

func (server restServer) doGetGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameInfo, err := server.controller.GetGame(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	gameJson, _ := json.Marshal(gameInfo)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gameJson)
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

func (server restServer) doNewGuess(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var guess data.Guess
	err = json.Unmarshal(body, &guess)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	params := mux.Vars(r)
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
