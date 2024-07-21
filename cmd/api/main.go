package main

import (
    "log"

	"github.com/NayanTheSpaceGuy/nayanpatil.space/v1/internal/server"
)

func main() {
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
