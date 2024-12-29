package main

import (
	"fmt"
	"github.com/veera-1729/lead-managment/internals/database"
	"github.com/veera-1729/lead-managment/internals/entities/contacts"
	"github.com/veera-1729/lead-managment/internals/entities/interactions"
	"github.com/veera-1729/lead-managment/internals/entities/leads"
	"github.com/veera-1729/lead-managment/internals/entities/orders"
	"github.com/veera-1729/lead-managment/internals/entities/restaurant"
	registry2 "github.com/veera-1729/lead-managment/internals/registry"
	"github.com/veera-1729/lead-managment/internals/router"
	contacts2 "github.com/veera-1729/lead-managment/internals/services/contacts"
	interactions2 "github.com/veera-1729/lead-managment/internals/services/interactions"
	"github.com/veera-1729/lead-managment/internals/services/lead"
	"github.com/veera-1729/lead-managment/internals/services/order"
	restaurant2 "github.com/veera-1729/lead-managment/internals/services/restaurant"
	"net/http"
)

func main() {

	db := database.New()
	err := db.Connect()
	if err != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Connected to database successfully")
	}

	userStore := contacts.NewUserStore(db)
	orderStore := orders.NewOrderStore(db)
	leadStore := leads.NewLeadStore(db)
	interactionStore := interactions.NewInteractionStore(db)
	restaurantStore := restaurant.NewRestaurantStore(db)

	userService := contacts2.NewContactService(userStore)
	restaurantService := restaurant2.NewRestaurantService(restaurantStore)
	leadService := lead.NewLeadService(leadStore)
	orderService := order.NewOrderService(orderStore)
	interactionService := interactions2.NewInteractionService(interactionStore)
	registry := registry2.GetRegistry()

	registry.RegisterServices(userService, restaurantService, leadService, orderService, interactionService)

	routeHandler := router.InitializeRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: routeHandler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic("Error starting server")
	} else {
		fmt.Println("Server started successfully")
	}
}
