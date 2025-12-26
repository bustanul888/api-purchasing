package item

import (
	"task-be/app/model"
)

type Service interface {
	create(req itemRequest) error
	getAll() []itemResponse
	getById(id string) itemResponse
	update(id string,req itemRequest) error
	delete(id string) error
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) create(req itemRequest) error{
	return s.repository.create(model.Item{
		Name: req.Name,
		Stock: req.Stock,
		Price: req.Price,
	})
}

func (s *service_) getAll() []itemResponse{
	return s.repository.getAll()
}

func (s *service_) getById(id string) itemResponse{
	return s.repository.GetById(id)
}

func (s *service_) update(id string,req itemRequest) error{
	return s.repository.update(id,req.Name,req.Stock,req.Price)
}

func (s *service_) delete(id string) error{
	return s.repository.delete(id)
}