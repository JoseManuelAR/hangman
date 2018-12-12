package view

import (
	"controller"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (server restServer) newGame(w http.ResponseWriter, r *http.Request) {
	game := server.controller.NewGame()
	w.Header().Set("Location", strings.Join([]string{r.Host, "games", game.Id}, "/"))
	w.WriteHeader(http.StatusNoContent)
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

func (server restServer) makeAGuess(w http.ResponseWriter, r *http.Request) {
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
	server.controller.MakeAGuess(params["id"], guess.Guess)
	w.WriteHeader(http.StatusNoContent)
}

type RestConfig struct {
	Ip   string
	Port string
}

type restServer struct {
	controller controller.Controller
	config     RestConfig
	router     *mux.Router
}

func NewRestServer(controller controller.Controller, config RestConfig) View {
	return &restServer{
		controller: controller,
		config:     config,
		router:     mux.NewRouter()}
}

func (server restServer) Start(wg sync.WaitGroup) error {
	log.Println("Starting rest server...")

	server.router.Use(commonMiddleware)
	server.router.HandleFunc("/games", server.getGamesInfo).Methods("GET")
	server.router.HandleFunc("/games", server.newGame).Methods("POST")
	server.router.HandleFunc("/games/{id}/guesses", server.makeAGuess).Methods("PUT")
	loggedRouter := handlers.LoggingHandler(os.Stdout, server.router)
	http.ListenAndServe(server.config.Ip+":"+server.config.Port, loggedRouter)
	return nil
}
