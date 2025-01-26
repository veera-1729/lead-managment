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

func (d *Db) Create(mode IModel) error {
	res := d.Db.Create(mode)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *Db) Find(id string, receiver IModel) error {
	res := d.Db.Where("id = ?", id).First(&receiver)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *Db) Update(model IModel) error {
	return nil
}

func (d *Db) Delete(id string) error {
	return nil
}
