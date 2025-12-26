package supplier

import (
	"fmt"
	"task-be/app/model"
)

type Service interface {
	create(req supplierRequest) error
	getAll() []supplierResponse
	update(id string,req supplierRequest) error
	delete(id string) error
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) create(req supplierRequest) error{
	fmt.Println("reqservice",req)
	return s.repository.create(model.Supplier{
		Name: req.Name,
		Email: req.Email,
		Address: req.Address,
	})
}

func (s *service_) getAll() []supplierResponse{
	return s.repository.getAll()
}

func (s *service_) getById(id string) supplierResponse{
	return s.repository.getById(id)
}

func (s *service_) update(id string,req supplierRequest) error{
	return s.repository.update(id,req.Name,req.Email,req.Address)
}

func (s *service_) delete(id string) error{
	return s.repository.delete(id)
}