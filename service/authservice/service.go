package authservice

import (
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/pooya-dehghan/entity"
)

type Config struct {
	SignKey                []byte
	Repo                   Repository
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
	AccessTokenSubject     string
	RefreshTokenSubject    string
}

type Service struct {
	config Config
}

func NewService(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.AccessTokenSubject, s.config.AccessTokenExpiration)

}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.RefreshTokenSubject, s.config.RefreshTokenExpiration)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
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

	return t.SignedString([]byte(s.config.SignKey))
}

type Repository interface {
	FindUserByPhoneNumber(phoneNumber string) (entity.User, error)
}
