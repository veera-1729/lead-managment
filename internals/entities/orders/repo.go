package orders

import "github.com/veera-1729/lead-managment/internals/database"

type Repo struct {
	DB *database.Db
}

func NewOrderStore(db *database.Db) *Repo {
	return &Repo{
		DB: db,
	}
}

// CreateOrder creates a new order
func (r *Repo) CreateOrder(order *Order) error {
	res := r.DB.Db.Create(order)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
