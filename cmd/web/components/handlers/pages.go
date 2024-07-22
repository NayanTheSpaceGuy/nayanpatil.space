package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/NayanTheSpaceGuy/nayanpatil.space/v1/cmd/web/templates/views"
	"github.com/NayanTheSpaceGuy/nayanpatil.space/v1/cmd/web/components/utils"
)

type IndexHandler struct {}

func(h IndexHandler) IndexDisplay(c echo.Context) error {
    return utils.Render(c, views.Base())
}
