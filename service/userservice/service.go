package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (createdUser entity.User, err error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	PhoneNumber string
	Name        string
	Password    string
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
		HashedPassword: GetMD5Hash(req.Password),
	}

	creatdUser, err := s.repo.RegisterUser(user)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("error in register user : %w", err)
	}

	return RegisterResponse{User: creatdUser}, nil
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	return LoginResponse{}, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
