package test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-rest-api-server/domain"
	"go-rest-api-server/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userPayload = `{"id": 1, "username": "John", "password": "password"}`

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assume you have a function GetUserByID that handles the request
	if assert.NoError(t, handler.GetUserByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		user := new(domain.User)
		err := json.Unmarshal(rec.Body.Bytes(), user)
		assert.NoError(t, err)

		// Verify the user details here
		//assert.Equal(t, 1, user.ID)
		//... other assertions
	}
}

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(userPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assume you have a function CreateUser that handles the request
	if assert.NoError(t, handler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		user := new(domain.User)
		err := json.Unmarshal(rec.Body.Bytes(), user)
		assert.NoError(t, err)

		// Verify the user details here
		assert.Equal(t, "John", user.Username)
		//... other assertions
	}
}
