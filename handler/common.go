package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func RootHandler(c echo.Context) error {
	tag, err := getLatestTag()
	if err != nil {
		return c.String(http.StatusOK, "go-punch-time@<error>")
	}
	return c.String(http.StatusOK, fmt.Sprintf("go-punch-time@%v", tag))
}

func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func getLatestTag() (string, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	tag := out.String()
	return tag[:len(tag)-1], nil
}
