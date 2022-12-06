package delivery

import (
	"be13/project/features/auth"
	"be13/project/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthDelivery struct {
	authServices auth.ServiceInterface
}

func NewAuth(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthDelivery{
		authServices: service,
	}

	e.POST("/login", handler.login)

}

func (delivery *AuthDelivery) login(c echo.Context) error {
	authInput := AuthRequest{}
	errBind := c.Bind(&authInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	token, hashed, err := delivery.authServices.Login(authInput.Email, authInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}

	z := []byte(authInput.Password)
	errPass := bcrypt.CompareHashAndPassword([]byte(hashed), z)
	if errPass != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Incorrect Password "+errPass.Error()))
	}

	data := map[string]any{
		"token": token,
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", data))

}