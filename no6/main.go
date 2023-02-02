package main

import (
	"log"

	"tog.test/no6/config"
	"tog.test/no6/model"
	"tog.test/no6/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.Connect()
	defer config.Close(db)

	db.AutoMigrate(&model.ShoppingCart{})

	//setup routes
	r := router.SetupRouter()

	// running
	r.Run(":3000")
}
