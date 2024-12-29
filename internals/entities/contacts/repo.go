package contacts

import "github.com/veera-1729/lead-managment/internals/database"

type Repo struct {
	DB *database.Db
}

func NewUserStore(db *database.Db) *Repo {
	return &Repo{
		DB: db,
	}
}

// CreateUser creates a new user
func (r *Repo) CreateUser(user *Contacts) error {
	res := r.DB.Db.Create(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *Repo) FetchAllUsers() ([]Contacts, error) {
	var users []Contacts
	res := r.DB.Db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}
