package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB(){
	str := "colocar credenciais do seu banco de dados postgres"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil{
		log.Fatal(err.Error())
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10) // conexoes maximas
	config.SetMaxOpenConns(100) // conexoes maximas abertas
	config.SetConnMaxLifetime(time.Hour) // tempo total da conexao (1 hora)
}

func GetDatabase() *gorm.DB{
	return db
}



