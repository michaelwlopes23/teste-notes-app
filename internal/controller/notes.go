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

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID do usuário inválido!")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "Erro ao salvar nota no banco de dados!")
		return
	}

	note.UserID = uint(userID)

	view.JSONResponse(w, http.StatusCreated, note)
}

func GetAllNote(w http.ResponseWriter, r *http.Request) {
	var notes []model.Note
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID do usuário inválido!")
		return
	}

	if err := database.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		view.ErrorResponse(w, http.StatusInternalServerError, "Erro ao buscar as notas do usuário!")
		return
	}

	view.JSONResponse(w, http.StatusOK, notes)
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID, err := strconv.Atoi(vars["id"])
	userID, err := strconv.Atoi(vars["userID"])

	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID invalido!")
		return
	}

	var note model.Note

	if err := database.DB.Where("node_id = ? user_id = ?", noteID, userID).First(&note).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
		return
	}

	view.JSONResponse(w, http.StatusOK, note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID, err := strconv.Atoi(vars["id"])
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var UpdateNote model.Note
	if err := json.NewDecoder(r.Body).Decode(&UpdateNote); err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "Erro ao processar o corpo da requisição!")
	}

	var existingNote model.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&existingNote).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
		return
	}

	existingNote.Title = UpdateNote.Title
	existingNote.Description = UpdateNote.Description

	if err := database.DB.Save(&existingNote).Error; err != nil {
		view.ErrorResponse(w, http.StatusInternalServerError, "Erro ao atualizar os dados de nota!")
	}

	view.JSONResponse(w, http.StatusOK, existingNote)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID, err := strconv.Atoi(vars["id"])
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID inválido!")
		return
	}

	var note model.Note

	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
		return
	}

	if err := database.DB.Delete(&note).Error; err != nil {
		view.ErrorResponse(w, http.StatusInternalServerError, "Erro ao deletar a nota!")
		return
	}

	view.JSONResponse(w, http.StatusOK, map[string]string{"message": "Nota deletada com sucesso!"})
}
