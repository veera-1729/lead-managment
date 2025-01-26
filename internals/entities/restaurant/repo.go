package restaurant

import (
	"github.com/veera-1729/lead-managment/internals/database"
)

type Repo struct {
	database.DataBase
}

func NewRestaurantStore(db database.DataBase) *Repo {
	return &Repo{
		db,
	}
}

// CreateRestaurant creates a new restaurant
func (r *Repo) CreateRestaurant(restaurant database.IModel) error {
	err := r.Create(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FetchAllRestaurants() ([]Restaurant, error) {
	//var restaurants []Restaurant
	//res := r.DB.Db.Find(&restaurants)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	//return restaurants, nil
	return nil, nil
}
