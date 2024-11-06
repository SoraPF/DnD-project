package router

import "github.com/gofiber/fiber/v2"

func FrontdRoutes() *fiber.App {
	router := fiber.New()
/*
	router.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("layouts/login", fiber.Map{
			"Title":      "LOGIN",
			"Categories": categories,
		})
	})

	router.Route("/authent", func(router fiber.Router) {
		router.Post("/login", Handlers.)
		router.Post("/register", Handlers.)
		router.Get("/isLogin", Handlers.)
		router.Post("/logout", Handlers.)
	})
*/
	return router
}
