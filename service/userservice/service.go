package service 

import (
	"pooyadehghan.com/entity"
	"pooyadehghan.com/phonenumber"
	_ "github.com/google/uuid"
)


type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool , error)
}

type Service struct {
	repo Repository
}
 
type  RegisterRequest struct{
	phoneNumber string
	name string
}
 
type  RegisterResponse struct{ 
	User entity.User
}

func (s Service) Register(req RegisterInput) (RegisterResponse , error){

	s.repo.IsPhoneNumberUnique(req.PhoneNumber)
}
