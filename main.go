package main

import (
	"notepad-api/configs"
	"notepad-api/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//loadEnv()
	configs.InitDB()
	e := echo.New()
	routes.InitRoutes(e)
	//e.Start(":8000")
	e.Logger.Fatal(e.Start(envPortOr("8000")))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal load .env")
	}
}

func envPortOr(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
