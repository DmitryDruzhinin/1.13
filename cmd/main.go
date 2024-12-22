package main

import (
	"github.com/DmitryDruzhinin/1.13/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
