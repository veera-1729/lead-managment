package restaurant

import "github.com/veera-1729/lead-managment/internals/database"

type Repo struct {
	DB *database.Db
}

func NewRestaurantStore(db *database.Db) *Repo {
	return &Repo{
		DB: db,
	}
}

// CreateRestaurant creates a new restaurant
func (r *Repo) CreateRestaurant(restaurant *Restaurant) error {
	res := r.DB.Db.Create(restaurant)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *Repo) FetchAllRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	res := r.DB.Db.Find(&restaurants)
	if res.Error != nil {
		return nil, res.Error
	}
	return restaurants, nil
}
