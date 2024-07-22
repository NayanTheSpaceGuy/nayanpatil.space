package main

import (
    "log"

	"github.com/NayanTheSpaceGuy/nayanpatil.space/v1/internal/server"
)

func main() {
    err := server.Init()
    if err != nil {
        log.Fatal(err)
    }
}
