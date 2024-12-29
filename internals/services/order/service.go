package order

import (
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/constants"
	"github.com/veera-1729/lead-managment/internals/entities/orders"
	"github.com/veera-1729/lead-managment/internals/registry"
)

type Service struct {
	repo *orders.Repo
}

func NewOrderService(repo *orders.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ServiceName() string {
	return constants.OrderService
}

func (s *Service) Handle(c *gin.Context) {
	panic("implement me")
}

func (s *Service) CreateOrder(c *gin.Context) {
	ord := &orders.Order{}
	err := c.ShouldBindBodyWithJSON(ord)
	if err != nil {
		c.JSON(400, gin.H{"error": "INVALID_REQUEST"})
		return
	}
	err = s.repo.CreateOrder(ord)
	if err != nil {
		c.JSON(500, gin.H{"error": "ERROR_CREATING_ORDER"})
		return
	}
	c.JSON(200, gin.H{"message": "ORDER_CREATED"})
}

func CreateOrder(c *gin.Context) {
	service := GetOrderServiceFromRegistry()
	service.CreateOrder(c)
}

func GetOrderServiceFromRegistry() *Service {
	r := registry.GetRegistry()
	var service *Service
	var ok bool
	if service, ok = r.GetServiceInterfaceFromRegistry(constants.OrderService).(*Service); !ok {
		panic("Error getting order service")
	}
	return service
}
