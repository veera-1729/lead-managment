package contacts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/constants"
	"github.com/veera-1729/lead-managment/internals/entities/contacts"
	"github.com/veera-1729/lead-managment/internals/registry"
	"net/http"
)

type Service struct {
	repo *contacts.Repo
}

func (s *Service) ServiceName() string {
	return constants.ContactService
}
func (s *Service) Handle(c *gin.Context) {
	fmt.Println(c.Request.URL.Path)
	switch c.Request.URL.Path {
	case "/user":
		//s.CreateContact(c)
	}
}

func NewContactService(repo *contacts.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func CreateContact(c *gin.Context) {
	service := GetContactServiceFromRegistry()
	user := &contacts.Contacts{}
	err := c.ShouldBindBodyWithJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
		return
	}
	err = service.repo.CreateContact(user)
	if err != nil {
		c.Set(constants.Response, nil)
		c.Set(constants.Error, err)
	} else {
		c.Set(constants.Response, user)
		c.Set(constants.Error, nil)
	}
}

func FetchAllContacts(c *gin.Context) {
	service := GetContactServiceFromRegistry()
	res, err := service.repo.FetchAllContacts()
	if err != nil {
		c.Set(constants.Response, nil)
		c.Set(constants.Error, err)
	} else {
		c.Set(constants.Response, res)
		c.Set(constants.Error, nil)
	}
}

func EditContact(c *gin.Context) {
	service := GetContactServiceFromRegistry()
	user := &contacts.Contacts{}
	err := c.ShouldBindBodyWithJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
		return
	}
	err = service.repo.EditContact(user, c.Param("id"))
	if err != nil {
		c.Set(constants.Response, nil)
		c.Set(constants.Error, err)
	} else {
		c.Set(constants.Response, user)
		c.Set(constants.Error, nil)
	}
}
func DeleteContact(c *gin.Context) {
	service := GetContactServiceFromRegistry()
	err := service.repo.DeleteContact(c.Param("id"))
	if err != nil {
		c.Set(constants.Response, nil)
		c.Set(constants.Error, err)
	} else {
		c.Set(constants.Response, "Deleted")
		c.Set(constants.Error, nil)
	}
}

func FetchRestaurantContacts(c *gin.Context) {
	service := GetContactServiceFromRegistry()
	res, err := service.repo.FetchContactsByRestaurantId(c.Param("id"))
	if err != nil {
		c.Set(constants.Response, nil)
		c.Set(constants.Error, err)
	} else {
		c.Set(constants.Response, res)
		c.Set(constants.Error, nil)
	}
}

func GetContactServiceFromRegistry() *Service {
	r := registry.GetRegistry()
	var service *Service
	var ok bool
	if service, ok = r.GetServiceInterfaceFromRegistry(constants.ContactService).(*Service); !ok {
		panic("Error getting user service")
	}
	return service
}
