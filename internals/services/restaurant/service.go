package restaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/constants"
	"github.com/veera-1729/lead-managment/internals/entities/restaurant"
	"github.com/veera-1729/lead-managment/internals/registry"
)

type Service struct {
	restaurantRepo *restaurant.Repo
}

func NewRestaurantService(repo *restaurant.Repo) *Service {
	return &Service{
		restaurantRepo: repo,
	}
}

func (s *Service) ServiceName() string {
	return constants.RestaurantService
}
func (s *Service) Handle(c *gin.Context) {
	switch c.Request.URL.Path {
	case "/restaurant":
		s.CreateRestaurant(c)
	}
}

func (s *Service) CreateRestaurant(c *gin.Context) (*restaurant.Restaurant, error) {
	re := &restaurant.Restaurant{}
	err := c.ShouldBindBodyWithJSON(re)
	if err != nil {
		return nil, err
	}
	err = s.restaurantRepo.CreateRestaurant(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}
func GetRestaurantServiceFromRegistry() *Service {
	r := registry.GetRegistry()
	var service *Service
	var ok bool
	if service, ok = r.GetServiceInterfaceFromRegistry(constants.RestaurantService).(*Service); !ok {
		panic("Error getting restaurant service")
	}
	return service
}

func CreateRestaurant(c *gin.Context) {
	service := GetRestaurantServiceFromRegistry()
	re, err := service.CreateRestaurant(c)
	if err != nil {
		c.Set(constants.Error, err)
		c.Set(constants.Response, nil)
	} else {
		c.Set(constants.Error, nil)
		c.Set(constants.Response, re)
	}
}

func FetchAllRestaurants(c *gin.Context) {
	service := GetRestaurantServiceFromRegistry()
	res, err := service.restaurantRepo.FetchAllRestaurants()
	if err != nil {
		c.Set(constants.Error, err)
		c.Set(constants.Response, nil)
	} else {
		c.Set(constants.Error, nil)
		c.Set(constants.Response, res)
	}
}
