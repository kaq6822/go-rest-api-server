package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-rest-api-server/domain"
	"go-rest-api-server/service"
	"net/http"
	"strconv"
)

func GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	res, err := service.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	res := service.CreateUser(*user)
	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	res := service.UpdateUser(user.ID, *user)
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	service.DeleteUser(id)

	return c.NoContent(http.StatusOK)
}

func Login(c echo.Context) error {
	login := new(domain.Login)
	if err := c.Bind(login); err != nil {
		fmt.Print(err)
		return err
	}

	fmt.Println(*login)

	isValid, err := service.Login(*login)
	if err != nil {
		return err
	}

	if !isValid {
		return c.NoContent(http.StatusUnauthorized)
	}

	jwt, err := service.NewJWT(login.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jwt)
}

func GetUser(c echo.Context) error {
	id := c.Get("id").(int)

	res, err := service.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
