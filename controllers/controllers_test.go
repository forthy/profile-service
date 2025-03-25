package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	E "github.com/IBM/fp-go/either"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func onCreateEcho() E.Either[error, *echo.Echo] {
	return E.Right[error](echo.New())
}

func onReleaseEcho(e *echo.Echo) E.Either[error, any] {
	if err := e.Shutdown(context.Background()); err != nil {
		return E.Left[any](err)
	} else {
		return E.Right[error, any]("Echo server is shutdown")
	}
}

const (
	sampleProfile  = `{"name":"John Doe","version":"1.0","domain":"example.com"}`
	sampleResponse = "User (name:John Doe, version:1.0, domain:example.com) was registered"
)

func TestRegister(t *testing.T) {
	// Method: POST
	// Path: /register

	E.WithResource[error, *echo.Echo, bool](
		onCreateEcho,
		onReleaseEcho,
	)(func(e *echo.Echo) E.Either[error, bool] {
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(sampleProfile))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, Register(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			return E.Right[error](assert.Equal(t, sampleResponse, rec.Body.String()))
		}

		return E.Right[error](false)
	})
}
