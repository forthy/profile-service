package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	CL "profile-svc/controllers"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	TC "github.com/xlab/closer"

	C "profile-svc/config"

	E "github.com/IBM/fp-go/either"
	ID "github.com/IBM/fp-go/identity"
	O "github.com/IBM/fp-go/option"
)

type FailedToLoadTcpPortVar struct {
}

func (b FailedToLoadTcpPortVar) Error() string {
	return "Failed to load TCP port variable"
}

func loadEnv(envFile string) error {
	return godotenv.Load(envFile)
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

var logger = log.Default()

// Context -> Echo -> ()
func cleanupFunc(e *echo.Echo) func() {
	return func() {
		logger.Print("Shutting down Echo...")

		if err := e.Shutdown(context.Background()); err != nil {
			logger.Fatal(err)
		}

		time.Sleep(3 * time.Second)

		logger.Print("  Done!")
	}
}

// Controller
/*
	Request -> Action -> Either error Profile
*/
func main() {
	// dot env
	go func() {
		e := echo.New()

		// Use xlab's closer: https://github.com/xlab/closer
		// to handle the graceful shutdown without using context.Context
		TC.Bind(cleanupFunc(e))

		e.Use(echoprometheus.NewMiddleware("profile_svc")) // adds middleware to gather metrics
		e.GET("/metrics", echoprometheus.NewHandler())     // adds route to serve gathered metrics

		e.GET("/", CL.HelloWorld)
		e.POST("/register", CL.Register)

		err := E.Fold(
			ID.Of[error],
			func(_ string) error {
				return O.Fold(
					func() error {
						return FailedToLoadTcpPortVar{}
					},
					func(p C.Port) error {
						return e.Start(fmt.Sprintf(":%d", p.Value))
					},
				)(C.ReadUnreservedPort("TCP_PORT"))
			},
		)(E.FromError(loadEnv)(".env"))

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	TC.Hold()
}
