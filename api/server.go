package api

import (
	"github.com/labstack/echo/v4"
	"go-rest-api-server/handler"
	"go-rest-api-server/middleware"
)

func InitAPI(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// user
	e.POST("/login", handler.Login)
	e.POST("/user", handler.CreateUser)

	g := e.Group("")
	g.Use(middleware.JWTMiddleware)
	g.GET("/user/:id", handler.GetUserByID)
	g.GET("/user", handler.GetUser)
	g.PUT("/user/:id", handler.UpdateUser)
	g.DELETE("/user/:id", handler.DeleteUser)
}
