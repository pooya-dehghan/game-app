package authservice

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/pkg/hash"
)

type Service struct {
	signKey                []byte
	repo                   Repository
	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
	accessTokenSubject     string
	refreshTokenSubject    string
}

func NewService(signKey string, accessTokenSubject string, refreshTokenSubject string, repo Repository, accessTokenExpiration time.Duration, refreshTokenExpiration time.Duration) Service {
	return Service{repo: repo, signKey: []byte(signKey), accessTokenExpiration: accessTokenExpiration, refreshTokenExpiration: refreshTokenExpiration, accessTokenSubject: accessTokenSubject, refreshTokenSubject: refreshTokenSubject}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.accessTokenSubject, s.accessTokenExpiration)

}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.refreshTokenSubject, s.refreshTokenExpiration)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims := token.Claims; token.Valid {
		return claims.(*Claims), nil
	} else {
		return nil, err
	}

}

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

func (c Claims) Valid() error {
	// return c.RegisteredClaims.Valid()
	return nil
}

func (s Service) createToken(userID uint, subject string, expiresAt time.Duration) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		UserID: userID,
	}

	return t.SignedString([]byte(s.signKey))
}

type Repository interface {
	FindUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
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

	token, err := s.createToken(user.ID, s.accessTokenSubject, s.accessTokenExpiration)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}

	return LoginResponse{AccessToken: token}, nil
}
