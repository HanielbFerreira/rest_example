package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hanbarfe/rest_example/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer o parse no ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error ao fazer o decode no JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.UpdateUserById(int64(id), user)
	if err != nil {
		log.Printf("Error ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Erro: foram atualizados %d registros ", rows)
	}

	res := map[string]any{
		"Erro":     false,
		"Mensagem": "Usuário atualizado com sucesso",
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)

}
