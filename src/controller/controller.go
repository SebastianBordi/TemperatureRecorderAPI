package controller

type Controller interface {
	GetAll() ([]struct{}, error)
	GetById(id int) (struct{}, error)
	Create(struct{}) error
	Delete(id int) error
}
