package auth

import (
	"errors"
	"fmt"
	"task-be/app/helper"
	"task-be/app/helper/blacklisttoken"
	"task-be/app/model"
	"task-be/app/service/user"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	login(request authRequest) (authResponse, int, error)
	logout(c *fiber.Ctx) error
}

type service_ struct {
	blackListRepo blacklisttoken.Repository
	userRepo user.Repository
}

func NewService(blackListRepo blacklisttoken.Repository,userRepo user.Repository) *service_ {
	return &service_{blackListRepo,userRepo}
}

func (s *service_) login(request authRequest) (authResponse, int, error) {
	user := s.userRepo.GetByUsername(request.UserName)
	if user.ID != "" {
		fmt.Println("User found")
		if helper.ComparePassword(request.Password, user.Password) {
			fmt.Println("Password matched")
			token, _ := helper.GenerateToken(user.ID, user.Role)
			return authResponse{
				ID:       user.ID,
				Token:    token,
				UserName: user.UserName,
				Role: user.Role,
				}, 200, nil
	}
	}
	return authResponse{}, 403, errors.New("incorrect username or password")
}

func (s *service_) logout(c *fiber.Ctx) error {
	token := helper.GetToken(c)
	s.blackListRepo.Create(model.BlackListToken{
		Token: *token,
	})
	return nil
}