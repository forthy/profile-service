package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	PM "profile-svc/model"

	E "github.com/IBM/fp-go/either"
	ID "github.com/IBM/fp-go/identity"
)

type RegisterBody struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Domain  string `json:"domain"`
}

type ActionId struct {
	Value string
}

type Action interface {
	ID() ActionId
}

type RegisterAction struct {
	Id ActionId
}

func (p *RegisterAction) ID() ActionId {
	return p.Id
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Helper function that helps Either.FromError function
func bindRegister(c echo.Context) func(b *RegisterBody) error {
	return func(b *RegisterBody) error {
		return c.Bind(b)
	}
}

// Register Profile -> Either error Profile -> echo.Context -> error
func Register(c echo.Context) error {
	return E.Fold(
		ID.Of[error],
		func(b *RegisterBody) error {
			// TODO - register the given profile
			_ = PM.DomainOf(b.Domain)

			return c.String(
				http.StatusCreated,
				fmt.Sprintf("User (name:%s, version:%s, domain:%s) was registered", b.Name, b.Version, b.Domain),
			)
		})(E.FromError(bindRegister(c))(new(RegisterBody)))
}
