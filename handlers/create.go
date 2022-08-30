package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hanbarfe/rest_example/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error ao fazer o decode no JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.CreateUser(user)

	var res map[string]any

	if err != nil {
		res = map[string]any{
			"Erro":     true,
			"Mensagem": fmt.Sprintf("Ocorreu um erro ao tentar criar o usuário: %v", err),
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res = map[string]any{
			"Erro":     false,
			"Mensagem": fmt.Sprintf("Usuário inserido com sucesso! Nome: %s, ID: %d", user.Name, id),
		}
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
