package main

import (
	"github.com/solairerove/go-simple-db-mux/app"
)

func main() {
	a := app.App{}
	a.Initialize("griotrard_user", "griotrard_pass", "griotrard")

	a.Run(":8000")
}
