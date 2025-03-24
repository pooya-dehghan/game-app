package authservice

import (
	"fmt"

	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/pkg/hash"
)

type Repository interface {
	FindUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	user, err := s.repo.FindUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("error in register user : %w", err)
	}

	if user.HashedPassword != hash.GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("password is not correct")
	}

	return LoginResponse{}, nil
}
