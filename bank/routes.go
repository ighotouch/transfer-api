package bank

import (
	"fmt"
	"net/http"
	"os"
	"transfer-api/bank/controllers"

	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "80"
	}

	e.Validator = &CustomValidator{validator: validator.New()}

	basePath := e.Group("/api/v1")

	transferController := new(controllers.TransferController)
	accountController := new(controllers.AccountController)

	basePath.POST("/transfer", transferController.InitiateTransfer)
	basePath.POST("/open-account", accountController.Create)

	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
