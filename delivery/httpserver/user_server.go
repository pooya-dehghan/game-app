package httpserver

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pooya-dehghan/service/userservice"
)

const SIGNED_KEY = "jwt_key"

func (s Server) userLogin(c echo.Context) error {
	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userService.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) userRegister(c echo.Context) error {
	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userService.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userProfile(c echo.Context) error {
	fmt.Println("c.GetAuthorization", c.Get("Authorization"))
	authToken := c.Request().Header.Get("Authorization")
	claims, err := s.authService.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userService.Profile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
