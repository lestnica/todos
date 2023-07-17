package main

import (
	"todo/pkg/service"
)

func main() {
	todoStart := service.NewService()

	todoStart.RunTodoCLI()
}
