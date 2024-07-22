package server

import (
    "github.com/NayanTheSpaceGuy/nayanpatil.space/v1/internal/server/env"
    "github.com/NayanTheSpaceGuy/nayanpatil.space/v1/internal/server/routes"
)

func Init() error {
    err := env.Init()
    if err != nil {
        return err
    }

    return routes.Init()
}
