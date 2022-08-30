package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hanbarfe/rest_example/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()

	if err != nil {
		log.Printf("Erro ao obter usuários: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer o parse no ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := models.GetUserById(int64(id))
	if err != nil {
		log.Printf("Error ao buscar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)

}
