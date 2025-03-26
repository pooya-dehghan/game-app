package authservice

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/pkg/hash"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
}

func createToken(userID uint, signedKey string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		UserID: userID,
	}

	return t.SignedString(signedKey)
}

type Repository interface {
	FindUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type Service struct {
	signKey string
	repo    Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	user, err := s.repo.FindUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("error in register user : %w", err)
	}

	if user.HashedPassword != hash.GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("password is not correct")
	}

	token, err := createToken(user.ID, s.signKey)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}

	return LoginResponse{AccessToken: token}, nil
}
