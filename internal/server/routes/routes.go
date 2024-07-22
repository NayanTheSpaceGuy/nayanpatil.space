package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/NayanTheSpaceGuy/nayanpatil.space/v1/cmd/web/components/handlers"
    "github.com/NayanTheSpaceGuy/nayanpatil.space/v1/internal/server/env"
)

func Init() error {
	e := echo.New()

    indexHandler := handlers.IndexHandler{}
    e.GET("/", indexHandler.IndexDisplay)

	return e.Start(":" + env.Port)
}
