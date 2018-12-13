package view

import (
	"config"
	"controller"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (server restServer) newGame(w http.ResponseWriter, r *http.Request) {
	gameInfo, err := server.controller.NewGame()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	gameInfoJson, _ := json.Marshal(gameInfo)
	w.WriteHeader(http.StatusOK)
	w.Write(gameInfoJson)
}

func (server restServer) getGamesInfo(w http.ResponseWriter, r *http.Request) {
	games := server.controller.GetGamesInfo()
	gamesJson, _ := json.Marshal(games)
	w.WriteHeader(http.StatusOK)
	w.Write(gamesJson)
}

type userGuess struct {
	Guess string
}

func (server restServer) newGuess(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var guess userGuess
	err = json.Unmarshal(body, &guess)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	gameInfo, err := server.controller.NewGuess(params["id"], guess.Guess)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	gameInfoJson, _ := json.Marshal(gameInfo)
	w.WriteHeader(http.StatusOK)
	w.Write(gameInfoJson)
}

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

	server.router.Use(commonMiddleware)
	server.router.HandleFunc("/games", server.getGamesInfo).Methods("GET")
	server.router.HandleFunc("/games", server.newGame).Methods("POST")
	server.router.HandleFunc("/games/{id}/guesses", server.newGuess).Methods("PUT")
	loggedRouter := handlers.LoggingHandler(os.Stdout, server.router)
	http.ListenAndServe(server.ip+":"+server.port, loggedRouter)
	return nil
}
