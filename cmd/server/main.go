package main

import (
	"log"
	"net/http"
	"notes-app/internal/controller"
	"notes-app/internal/database"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", controller.CreateUser).Methods("POST")

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/notes", controller.CreateNote).Methods("POST")
	r.HandleFunc("/notes", controller.GetAllNote).Methods("GET")
	r.HandleFunc("/notes/{id}", controller.GetNoteByID).Methods("GET")
	r.HandleFunc("/notes/{id}", controller.UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/{id}", controller.DeleteNote).Methods("DELETE")

	log.Println("Iniciando o servidor na porta 8000...")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo ao aplicativo de notas!"))
}
