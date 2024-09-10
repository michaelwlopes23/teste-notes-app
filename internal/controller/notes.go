package controller

import (
	"encoding/json"
	"net/http"
	"notes-app/internal/model"
	"notes-app/internal/view"
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
