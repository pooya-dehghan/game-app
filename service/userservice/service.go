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
	signedKey []byte
	repo      Repository
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User UserResponse
}

type UserResponse struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}

func New(repo Repository, signedKey []byte) Service {
	return Service{repo: repo, signedKey: signedKey}
}

func (s Service) Register(req RegisterRequest) (UserResponse, error) {
	if !phonenumber.IsValid(req.PhoneNumber) {
		return UserResponse{}, fmt.Errorf("phone number is not valid")
	}

	if req.Name == "" {
		return UserResponse{}, fmt.Errorf("name is empty")
	}

	if len(req.Password) < 8 {
		return UserResponse{}, fmt.Errorf("password is not long enough")
	}

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {

		if err != nil {
			return UserResponse{}, err
		}

		if !isUnique {
			return UserResponse{}, fmt.Errorf("phone number is not unique")
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
		return UserResponse{}, fmt.Errorf("error in register user : %w", err)
	}

	userResponse := UserResponse{
		ID:          creatdUser.ID,
		PhoneNumber: creatdUser.PhoneNumber,
		Name:        creatdUser.Name,
	}

	return userResponse, nil
}

type ProfileRequest struct {
	UserID uint
}

func (s Service) Profile(req ProfileRequest) (entity.User, error) {

	user, err := s.repo.FindUserByID(req.UserID)

	if err != nil {
		fmt.Println(err)
	}

	return user, nil
}
