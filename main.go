package main

import (
	"fmt"

	"github.com/IqbalLx/alterra-agmc/config"
	"github.com/IqbalLx/alterra-agmc/utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	e.Validator = (utils.NewRequestValidator(validator.New()))

	v1 := e.Group("/v1")
	InitializeServer(v1)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", config.Server.Port)))
}
