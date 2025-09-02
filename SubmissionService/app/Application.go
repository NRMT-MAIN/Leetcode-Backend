package app

import (
	"Submission_Service/config/env"
	"fmt"
	"net/http"
	"time"
)


type Config struct {
	Address string
}

type Application struct {
	Config Config
}

func NewConfig() *Config {
	PORT := env.GetString("PORT" , ":3000")

	return &Config {
		Address: PORT,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}


func (app *Application) Run() error {
	server := &http.Server{
		Addr: app.Config.Address,
		Handler: nil,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Println("Server is running on port :" , app.Config.Address) ;
	return server.ListenAndServe() ; 
}