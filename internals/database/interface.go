package database

type IModel interface {
	TableName() string
}

type Creator interface {
	Create(data IModel) error
}

type Updater interface {
	Update(data IModel) error
}

type Finder interface {
	Find(id string, receiver IModel) error
}

type Deleter interface {
	Delete(id string) error
}

type DataBase interface {
	Creator
	Updater
	Finder
	Deleter
}
