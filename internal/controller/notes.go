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

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "Dados inválidos!")
		return
	}

	if err := database.DB.Create(&note).Error; err != nil {
		view.ErrorResponse(w, http.StatusInternalServerError, "Erro ao saver a nota no banco de dados!")
	}

	view.JSONReponse(w, http.StatusCreated, note)
}

func GetAllNote(w http.ResponseWriter, r *http.Request) {
	var notes []model.Note

	if err := database.DB.Find(&notes).Error; err != nil {
		view.ErrorResponse(w, http.StatusInternalServerError, "Erro ao buscar as notas")
	}

	view.JSONReponse(w, http.StatusOK, notes)
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID invalido!")
		return
	}

	var note model.Note
	if err := database.DB.First(&note, id).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
	}

	view.JSONReponse(w, http.StatusOK, note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID inválido!")
		return
	}

	var note model.Note

	if err := database.DB.First(&note, id).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
		return
	}

	if err := database.DB.Delete(&note, id).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Erro ao deletar a nota!")
		return
	}

	view.JSONReponse(w, http.StatusOK, map[string]string{"message": "Nota deletada com sucesso!"})
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var UpdateNote model.Note

	if err := json.NewDecoder(r.Body).Decode(&UpdateNote); err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "Erro ao processar o corpo da requisição!")
	}

	var existingNote model.Note

	if err := database.DB.First(&existingNote, id).Error; err != nil {
		view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
	}

	existingNote.Content = UpdateNote.Content

}
