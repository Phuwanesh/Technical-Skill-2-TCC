package main

import (
	"fmt"
	"log"

	config "authdemo/database"
	"authdemo/handlers"
	"authdemo/models"
	"authdemo/routes"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("warning: no .env file found or failed to load")
	}

	config.InitDB()
	db := config.GetDB()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}
	store := &handlers.GormUserStore{DB: db}
	r := routes.SetupRouter(store)

	log.Println("API listening at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
