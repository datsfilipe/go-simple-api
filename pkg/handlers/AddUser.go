package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"math/rand"
	"io"

	"github.com/datsfilipe/go-simple-api/pkg/models"
	"github.com/datsfilipe/go-simple-api/pkg/mocks"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	user.Id = rand.Intn(1000000)
	mocks.Users = append(mocks.Users, user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created!")
}
