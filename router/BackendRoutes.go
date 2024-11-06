package router

import "github.com/gofiber/fiber/v2"

func BackendRoutes() *fiber.App {
	router := fiber.New()
/*
	router.Route("/authent", func(router fiber.Router) {
		router.Post("/login", Handlers.)
		router.Post("/register", Handlers.)
		router.Get("/isLogin", Handlers.)
		router.Post("/logout", Handlers.)
	})
*/
	return router
}
