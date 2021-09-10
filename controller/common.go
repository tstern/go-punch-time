package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type CommonController struct {
	tag string
}

func NewCommonController() *CommonController {
	var tag string
	var ok bool

	if tag, ok = os.LookupEnv("PUNCH_TIME_VERSION"); !ok {
		tag = "unknown"
	}

	return &CommonController{
		tag: tag,
	}
}

func (ctrl *CommonController) Root(e echo.Context) error {
	return e.String(http.StatusOK, fmt.Sprintf("go-punch-time@%v", ctrl.tag))
}

func (ctrl *CommonController) Pong(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}
