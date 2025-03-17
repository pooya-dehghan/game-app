package service

import (
	"fmt"

	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	phoneNumber string
	name        string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	if !phonenumber.IsValid(req.phoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.phoneNumber); err != nil || !isUnique {

		if err != nil {
			return RegisterResponse{}, err
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	return RegisterResponse{}, nil
}
