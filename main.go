package main

import (
	"github.com/sobatfillah/config"
	"github.com/sobatfillah/migration"
	"github.com/sobatfillah/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
	db := config.Init(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	migration.Migrate(db)
}

func main(){
	gin.SetMode(gin.ReleaseMode)

	router := route.Routes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	if err := router.Run(":" + port);err != nil {
		log.Panicf("error: %s", err)
	}
}