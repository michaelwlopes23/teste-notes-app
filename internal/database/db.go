package database

import (
	"log"
	"notes-app/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados", err)
	}

	err = DB.AutoMigrate(&model.Note{})
	if err != nil {
		log.Fatal("Falha ao migrar a tabela:", err)
	}

	log.Println("Banco de dados conectado e migrado com sucesso!")
}
