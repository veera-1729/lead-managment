package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Db *gorm.DB
}

func New() *Db {
	return &Db{}
}

func (d *Db) Connect() error {
	info := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		"localhost",
		"veera",
		"password",
		"lead_managment",
		"5433",
	)

	db, err := gorm.Open(postgres.Open(info), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	d.Db = db
	return nil
}
