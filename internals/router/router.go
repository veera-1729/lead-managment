package router

import (
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/middleware"
	"github.com/veera-1729/lead-managment/internals/services/contacts"
	"github.com/veera-1729/lead-managment/internals/services/interactions"
	"github.com/veera-1729/lead-managment/internals/services/lead"
	"github.com/veera-1729/lead-managment/internals/services/order"
	"github.com/veera-1729/lead-managment/internals/services/restaurant"
)

func InitializeRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to User API",
		})
	})

	router.Use(middleware.ResponseHandler())

	setUpUserRoutes(router)
	setUpRestaurantRoutes(router)
	setUpLeadRoutes(router)
	setUpOrderRoutes(router)
	setUpInteractionRoutes(router)
	return router
}

func setUpUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/contacts")

	userRoutes.GET("", contacts.FetchAllContacts)
	userRoutes.POST("", contacts.CreateContact)
	//userRoutes.PUT("/user/:id", UpdateUser)
	//userRoutes.DELETE("/user/:id", DeleteUser)
}
func setUpRestaurantRoutes(router *gin.Engine) {
	restaurantRoutes := router.Group("/restaurant")

	restaurantRoutes.GET("", restaurant.FetchAllRestaurants)
	restaurantRoutes.POST("", restaurant.CreateRestaurant)
	//restaurantRoutes.PUT("/restaurant/:id", UpdateRestaurant)
	//restaurantRoutes.DELETE("/restaurant/:id", DeleteRestaurant)
}

func setUpLeadRoutes(router *gin.Engine) {
	leadRoutes := router.Group("/lead")

	leadRoutes.GET("", lead.FetchAllLeads)
	leadRoutes.POST("", lead.CreateLead)
	//leadRoutes.PUT("/lead/:id", UpdateLead)
	//leadRoutes.DELETE("/lead/:id", DeleteLead)
}

func setUpOrderRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/order")

	//orderRoutes.GET("/order/:id", GetOrder)
	orderRoutes.POST("", order.CreateOrder)
	//orderRoutes.PUT("/order/:id", UpdateOrder)
	//orderRoutes.DELETE("/order/:id", DeleteOrder)
}

func setUpInteractionRoutes(router *gin.Engine) {
	interactionRoutes := router.Group("/interaction")

	//interactionRoutes.GET("/interaction/:id", GetInteraction)
	interactionRoutes.POST("", interactions.CreateInteraction)
	//interactionRoutes.PUT("/interaction/:id", UpdateInteraction)
	//interactionRoutes.DELETE("/interaction/:id", DeleteInteraction)
}
