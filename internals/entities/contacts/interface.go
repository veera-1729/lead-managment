package contacts

import "github.com/veera-1729/lead-managment/internals/database"

type Creator interface {
	CreateContact(data database.IModel) error
}

type Finder interface {
	FindByID(id string, receiver database.IModel) error
}
