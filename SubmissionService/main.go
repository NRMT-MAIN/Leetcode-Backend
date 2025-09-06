package main

import (
	"Submission_Service/app"
	"Submission_Service/config/env"
)

func main() {

	env.Load() ; 
	cfg := app.NewConfig() 

	app := app.NewApplication(cfg) ; 

	app.Run()
}