package app

import (
	config "Submission_Service/config/db"
	"Submission_Service/config/env"
	"Submission_Service/controllers"
	"Submission_Service/db"
	"Submission_Service/db/repositories"
	"Submission_Service/routers"
	"Submission_Service/service"
	"fmt"
	"net/http"
	"time"
)


type Config struct {
	Address string
	Store   *db.Storage
}

type Application struct {
	Config Config
}

func NewConfig() *Config {
	PORT := env.GetString("PORT" , ":3000")

	return &Config {
		Address: PORT,
		Store: db.NewStorage(),
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}


func (app *Application) Run() error {
	dbClient  , err := config.CreateClient()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return err
	}	

	pr := repositories.NewSubmissionRepository(dbClient)
	ps := service.NewSubmissionService(pr) ; 
	pc := controllers.NewSubmissionController(ps) ; 
	prt := routers.NewSubmissionRouter(pc) ; 

	server := &http.Server{
		Addr: app.Config.Address,
		Handler: routers.SetupRouter(prt),
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Println("Server is running on port :" , app.Config.Address) ;
	return server.ListenAndServe() ; 
}