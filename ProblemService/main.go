package main

import (
	"fmt"
	"leetcode/app"
	config "leetcode/config/db"
)

func main() {
    // Your code here
	config.CreateClient()
	cfg := app.NewConfig()
	application := app.NewApplication(cfg)
	
	if err := application.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}