package main

import (
	"log"
	"net/http"
	"notes-app/internal/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/notes", controller.CreateNote).Methods("POST")

	log.Println("Iniciando o servidor na porta 8000...")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo ao aplicativo de notas!"))
}
