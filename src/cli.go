package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func customNewGame(wordsFile string) http.HandlerFunc {
	//words := hangman.ReadWordsFromFile(wordsFile)

	return func(w http.ResponseWriter, r *http.Request) {
		//choosenWord := hangman.PickWord(words)
		//game := hangman.NewGame(3, choosenWord)
		//store.CreateGame(game)
		//w.Header().Set("Location", strings.Join([]string{r.Host, "games", game.ID}, "/"))
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	var wordsFile_param string
	var hints_param string

	flag.StringVar(&wordsFile_param, "words_file", "words/words.txt", "Words file")
	flag.StringVar(&hints_param, "hints", "3", "Max number of hints")
	flag.Parse()

	hints, err := strconv.Atoi(hints_param)

	if err == nil {
		log.Println(wordsFile_param, hints)
	}

	router := mux.NewRouter()
	router.Use(commonMiddleware)
	// Register HTTP endpoints
	router.HandleFunc("/games", customNewGame(wordsFile_param)).Methods("POST")
	// router.HandleFunc("/games/{id}", retrieveGameInfo).Methods("GET")
	// router.HandleFunc("/games/{id}/guesses", makeAGuess).Methods("PUT")
	// router.HandleFunc("/games/{id}", deleteGame).Methods("DELETE")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
