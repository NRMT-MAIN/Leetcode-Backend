package app

import (
	"fmt"
	dbConfig "leetcode/config/db"
	"leetcode/config/env"
	"leetcode/controllers"
	"leetcode/db"
	"leetcode/db/repositories"
	"leetcode/routers"
	"leetcode/services"
	"net/http"
	"time"
)

type Config struct {
	Addr  string
	Store *db.Storage
}

type Application struct {
	Config *Config
}

func NewConfig() *Config {
	port := env.GetString("PORT", ":8080")

	return &Config{
		Addr:  port,
		Store: db.NewStorage(),
	}
}

func NewApplication(config *Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {

	client , _:= dbConfig.CreateClient() 
	pr := repositories.NewProblemRepository(client)
	ps := services.NewProblemService(pr)
	pc := controllers.NewProblemController(ps)
	prt := routers.NewProblemRouter(pc)


	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      routers.SetupRouter(prt),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}
