package service

import (
	"fmt"

	"pooyadehghan.com/entity"
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

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.phoneNumber); err != nil || isUnique != nil {

		if err != nil {
			return RegisterResponse{}, err
		}

		if isUnique != nil {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

}
