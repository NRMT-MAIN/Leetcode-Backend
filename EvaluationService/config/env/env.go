package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error in Loading environment variables") ; 
	}
}

func GetString(key string , fallback string) string {
	value , ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	return value
}

func GetInt(key string , fallback int) int {
	value , ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	result , err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("Error in converting the value")  ; 
		return fallback
	}

	return result
}