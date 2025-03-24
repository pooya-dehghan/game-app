package userservice

import (
	"fmt"

	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/pkg/hash"
	"github.com/pooya-dehghan/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (createdUser entity.User, err error)
	FindUserByID(userID uint) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password is not long enough")
	}

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {

		if err != nil {
			return RegisterResponse{}, err
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	user := entity.User{
		ID:             0,
		PhoneNumber:    req.PhoneNumber,
		Name:           req.Name,
		HashedPassword: hash.GetMD5Hash(req.Password),
	}

	creatdUser, err := s.repo.RegisterUser(user)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("error in register user : %w", err)
	}

	return RegisterResponse{User: creatdUser}, nil
}

type ProfileRequest struct {
	userID uint
}

func (s Service) Profile(req ProfileRequest) (entity.User, error) {

	user, err := s.repo.FindUserByID(req.userID)

	if err != nil {
		fmt.Println(err)
	}

	return user, nil
}
