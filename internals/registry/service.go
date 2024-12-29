package registry

import (
	"fmt"
	"sync"
)

var Once sync.Once
var serviceRegistry *Registry

func GetRegistry() *Registry {
	Once.Do(func() {
		serviceRegistry = NewRegistry()
	})
	return serviceRegistry
}

type Registry struct {
	register map[string]IService
}

func NewRegistry() *Registry {
	return &Registry{
		register: make(map[string]IService),
	}
}

func (r *Registry) RegisterServices(service ...IService) {
	for _, s := range service {
		r.register[s.ServiceName()] = s
	}
}

func (r *Registry) GetServiceInterfaceFromRegistry(name string) IService {
	if _, ok := r.register[name]; !ok {
		fmt.Printf("service %s not found", name)
		return nil
	}
	return r.register[name]
}

func (r *Registry) ClearRegistry() {
	r.register = map[string]IService{}
}
