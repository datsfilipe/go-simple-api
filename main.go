package main

import (
	"log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
)

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println("API running!")
	http.ListenAndServe(":4000", router)
}
