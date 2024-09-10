package controller

import (
	"encoding/json"
	"net/http"
	"notes-app/internal/model"
	"notes-app/internal/view"
	"strconv"

	"github.com/gorilla/mux"
)

var notes []model.Note
var lastID int

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote model.Note

	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		view.ErrorResponse(w, http.StatusBadRequest, "Dados inválidos!")
		return
	}

	lastID++
	newNote.ID = lastID
	notes = append(notes, newNote)

	view.JSONReponse(w, http.StatusCreated, newNote)
}

func GetAllNote(w http.ResponseWriter, r *http.Request) {
	if len(notes) == 0 {
		view.JSONReponse(w, http.StatusOK, "Nenhuma nota criada ainda!")
		return
	}

	view.JSONReponse(w, http.StatusOK, notes)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		View.ErrorResponse(w, http.StatusBadRequest, "ID inválido!")
		return
	}

	for index, note := range notes {
		if note.ID == id {
			notes = append(notes[:index], notes[index+1:]...)
			view.JSONReponse(w, http.StatusOK, "Nota deletada com sucesso!")
			return
		}
	}

	view.ErrorResponse(w, http.StatusNotFound, "Nota não encontrada!")
}
