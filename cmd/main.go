package main

import (
	"github.com/virtualmachinerydev/calc_go/internal/application"
)

func main() {
	app := application.New()
	// app.Run()
	app.RunServer()
}
