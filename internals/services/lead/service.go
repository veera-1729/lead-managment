package lead

import (
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/constants"
	"github.com/veera-1729/lead-managment/internals/entities/leads"
	"github.com/veera-1729/lead-managment/internals/registry"
)

type Service struct {
	leadRepo *leads.Repo
}

func NewLeadService(repo *leads.Repo) *Service {
	return &Service{
		leadRepo: repo,
	}
}

func (s *Service) ServiceName() string {
	return constants.LeadService
}

func (s *Service) Handle(c *gin.Context) {
	switch c.Request.URL.Path {
	case "/lead":
		s.CreateLead(c)
	}
}

func (s *Service) CreateLead(c *gin.Context) (*leads.Lead, error) {
	le := &leads.Lead{}
	err := c.ShouldBindBodyWithJSON(le)
	if err != nil {
		return nil, err
	}
	err = s.leadRepo.CreateLead(le)
	if err != nil {
		return nil, err
	}
	return le, nil
}

func CreateLead(c *gin.Context) {
	service := GetLeadServiceFromRegistry()
	s, err := service.CreateLead(c)
	if err != nil {
		c.Set(constants.Error, err)
		c.Set(constants.Response, nil)
	} else {
		c.Set(constants.Error, nil)
		c.Set(constants.Response, s)
	}
}

func FetchAllLeads(c *gin.Context) {
	service := GetLeadServiceFromRegistry()
	leads, err := service.leadRepo.FetchAllLeads()
	if err != nil {
		c.Set(constants.Error, err)
		c.Set(constants.Response, nil)
	} else {
		c.Set(constants.Error, nil)
		c.Set(constants.Response, leads)
	}
}

func GetLeadServiceFromRegistry() *Service {
	r := registry.GetRegistry()
	var service *Service
	var ok bool
	if service, ok = r.GetServiceInterfaceFromRegistry(constants.LeadService).(*Service); !ok {
		panic("Error getting lead service")
	}
	return service
}
