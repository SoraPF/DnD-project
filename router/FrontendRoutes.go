package router

import "github.com/gofiber/fiber/v2"

func FrontdRoutes() *fiber.App {
	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title":      "main",
		})
	})

	router.Get("/races", func(c *fiber.Ctx) error {
		return c.Render("pages/races", fiber.Map{
			"Title":      "races",
		})
	})

	router.Get("/background", func(c *fiber.Ctx) error {
		return c.Render("pages/background", fiber.Map{
			"Title":      "background",
		})
	})

	router.Get("/level",func (c* fiber.Ctx) error {
		return c.Render("pages/level",fiber.Map{
			"Title": "level",
		})
	})

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
