package interactions

import "github.com/veera-1729/lead-managment/internals/database"

type Repo struct {
	DB *database.Db
}

func NewInteractionStore(db *database.Db) *Repo {
	return &Repo{
		DB: db,
	}
}

// CreateInteraction  creates a new interaction
func (r *Repo) CreateInteraction(in *Interaction) error {
	res := r.DB.Db.Create(in)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
