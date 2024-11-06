package main

import (
	"DnD-project/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
    fmt.Print("Run service ...")

	loagConfig, err := config.LoadConfig("./env")
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
    //init database
	db := config.ConnectionDB(&loagConfig)
	err = config.AutoMigrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Database migration successful")

    // initialisation du moteur de templates
    fmt.Println("init templates")

    engine := django.New("./view",".django")
    engine.Reload(true)
    engine.Debug(true)

    //engine.AddFunc()

    fmt.Print("Run application")

    app := fiber.New(fiber.Config{
            Views: engine,
    })

    app.Static("/public","./public")

    // routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber!")
    })

    log.Fatal(app.Listen(":3000"))
}