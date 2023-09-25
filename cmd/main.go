package main

import (
	"log"
    "net/http"

    "github.com/gorilla/mux"
	"github.com/datsfilipe/go-simple-api/pkg/handlers"
)

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetAllUsers).Methods(http.MethodGet)

	log.Println("API running!")
	http.ListenAndServe(":4000", router)
}
