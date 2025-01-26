package contacts

import (
	"github.com/veera-1729/lead-managment/internals/database"
)

type Repo struct {
	database.DataBase
}

func NewUserStore(db database.DataBase) *Repo {
	return &Repo{
		db,
	}
}

// CreateContact  creates a new user
func (r *Repo) CreateContact(user database.IModel) error {
	//res := r.DB.Db.Create(user)
	//if res.Error != nil {
	//	return res.Error
	//}
	err := r.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FindByID(id string, receiver database.IModel) error {
	err := r.Find(id, receiver)
	if err != nil {
		return err
	}
	return nil
}

//func (r *Repo) Update(user database.IModel) error {
//	return nil
//}
//
//func (r *Repo) Delete(id string) error {
//	return nil
//}

//func (r *Repo) FetchAllContacts() ([]Contacts, error) {
//	var users []Contacts
//	res := r.DB.Db.Find(&users)
//	if res.Error != nil {
//		return nil, res.Error
//	}
//	return users, nil
//}
//
//func (r *Repo) EditContact(user *Contacts, id string) error {
//	res := r.DB.Db.Model(&Contacts{}).Where("id = ?", id).Updates(user)
//	if res.Error != nil {
//		return res.Error
//	}
//	return nil
//}
//
//func (r *Repo) DeleteContact(id string) error {
//	res := r.DB.Db.Where("id = ?", id).Delete(&Contacts{})
//	if res.Error != nil {
//		return res.Error
//	}
//	return nil
//}
//
//func (r *Repo) FetchContactsByRestaurantId(id string) ([]Contacts, error) {
//	var contacts []Contacts
//	res := r.DB.Db.Where("restaurant_id = ?", id).Find(&contacts)
//	if res.Error != nil {
//		return nil, res.Error
//	}
//	return contacts, nil
//}
