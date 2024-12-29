package leads

import "github.com/veera-1729/lead-managment/internals/database"

type Repo struct {
	DB *database.Db
}

func NewLeadStore(db *database.Db) *Repo {
	return &Repo{
		DB: db,
	}
}

// CreateLead creates a new lead
func (r *Repo) CreateLead(lead *Lead) error {
	res := r.DB.Db.Create(lead)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *Repo) FetchAllLeads() ([]Lead, error) {
	var leads []Lead
	res := r.DB.Db.Find(&leads)
	if res.Error != nil {
		return nil, res.Error
	}
	return leads, nil
}

func (r *Repo) EditLead(lead *Lead, id string) error {
	res := r.DB.Db.Model(&Lead{}).Where("id = ?", id).Updates(lead)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
