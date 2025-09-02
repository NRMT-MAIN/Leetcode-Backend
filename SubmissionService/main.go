package main

import (
	"Submission_Service/app"
)

func main() {
	cfg := app.NewConfig() 

	app := app.NewApplication(*cfg) ; 

	app.Run()
}