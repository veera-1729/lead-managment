package restaurant

import "github.com/veera-1729/lead-managment/internals/database"

type Creator interface {
	CreateRestaurant(model database.IModel) error
}

type Finder interface {
	FindByID(id string, receiver database.IModel) error
}
