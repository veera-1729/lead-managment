package main

import (
	"fmt"
	"github.com/veera-1729/lead-managment/internals/entities/contacts"
	"github.com/veera-1729/lead-managment/internals/entities/leads"
	"github.com/veera-1729/lead-managment/internals/entities/orders"
	"github.com/veera-1729/lead-managment/internals/entities/restaurant"
	"github.com/veera-1729/lead-managment/internals/services/interactions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	info := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s",
		"localhost",
		"veera",
		"password",
		"lead_managment",
		"5433",
		"public",
	)

	db, err := gorm.Open(postgres.Open(info), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Run migrations
	err = db.AutoMigrate(&contacts.Contacts{}, &restaurant.Restaurant{}, &leads.Lead{}, &orders.Order{}, &interactions.Service{})
	if err != nil {
		panic("failed to migrate database")
	}

	println("Database migrated successfully!")
}
