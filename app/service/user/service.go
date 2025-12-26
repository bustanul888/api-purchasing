package user

import (
	"errors"
	"task-be/app/helper"
	"task-be/app/model"
)

type Service interface {
	create(req userRequest) error
	getAll() []userResponse
	update(id string,req userUpdateRequest) error
	updatePassword(id string, req userUpdatePassword) (int,error)
	delete(id string) error
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) create(req userRequest) error{
	password := helper.HashPasword(req.Password)
	return s.repository.create(model.User{
		UserName: req.UserName,
		Password: password,
		Role: req.Role,
	})
}

func (s *service_) getAll() []userResponse{
	return s.repository.getAll()
}

func (s *service_) getById(id string) userResponse{
	return s.repository.GetById(id)
}

func (s *service_) update(id string,req userUpdateRequest) error{
	return s.repository.update(id,req.UserName,req.Role)
}

func (s *service_) updatePassword(id string, req userUpdatePassword) (int,error){
	user := s.repository.GetById(id)
	if !helper.ComparePassword(req.OldPassword, user.Password) {
		return 409, errors.New("OLD_PASSWORD_FALSE")
	}
	newHashPassword := helper.HashPasword(req.NewPassword)
	err := s.repository.updatePassword(id,newHashPassword)
	if err!=nil{
		return 500,err
	}
	return 200,err
}

func (s *service_) delete(id string) error{
	return s.repository.delete(id)
}