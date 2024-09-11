package controller

import (
	"encoding/json"
	"net/http"
	"notes-app/internal/database"
	"notes-app/internal/model"
	"notes-app/internal/view"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "Erro ao processar o corpo da requisição")
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		view.ErrorResponse(w, http.StatusInternalServerError, "Erro ao salvar o usuário no banco de dados")
		return
	}

	view.JSONResponse(w, http.StatusCreated, user)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var user model.User

	if err := database.DB.First(&user, id).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Usuário não encontrado")
		return
	}

	view.JSONResponse(w, http.StatusOK, user)
}
