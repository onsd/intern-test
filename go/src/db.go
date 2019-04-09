package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		"host="+os.Getenv("HOSTNAME")+
			"port="+os.Getenv("PORT")+
			"user="+os.Getenv("USER")+
			"dbname="+os.Getenv("DBNAME")+
			"password="+os.Getenv("PASSWORD"),
	)
	if err != nil {
		log.Printf("Error at open Database: %v", err)
		return nil, err
	}
	return db, nil
}
