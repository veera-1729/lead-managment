package interactions

import (
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/constants"
	"github.com/veera-1729/lead-managment/internals/entities/interactions"
	"github.com/veera-1729/lead-managment/internals/registry"
)

type Service struct {
	irepo *interactions.Repo
}

func NewInteractionService(repo *interactions.Repo) *Service {
	return &Service{
		irepo: repo,
	}
}

func (s *Service) ServiceName() string {
	return constants.InteractionsService
}

func (s *Service) Handle(c *gin.Context) {
	panic("implement me")
}

func (s *Service) CreateInteraction(c *gin.Context) {
	i := &interactions.Interaction{}
	err := c.ShouldBindBodyWithJSON(i)
	if err != nil {
		c.JSON(400, gin.H{"error": "INVALID_REQUEST"})
		return
	}
	err = s.irepo.CreateInteraction(i)
	if err != nil {
		c.JSON(500, gin.H{"error": "ERROR_CREATING_INTERACTION"})
		return
	}
	c.JSON(200, gin.H{"message": "INTERACTION_CREATED"})
}

func CreateInteraction(c *gin.Context) {
	service := GetInteractionServiceFromRegistry()
	service.CreateInteraction(c)
}
func GetInteractionServiceFromRegistry() *Service {
	r := registry.GetRegistry()
	var service *Service
	var ok bool
	if service, ok = r.GetServiceInterfaceFromRegistry(constants.InteractionsService).(*Service); !ok {
		panic("Error getting restaurant service")
	}
	return service
}
