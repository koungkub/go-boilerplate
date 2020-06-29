package route

import (
	"boilerplate/internal/handler"
	"boilerplate/internal/model"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"
)

func routing(app *fiber.App) {
	app.Get("/", handler.GetReview())
}

// New get fiber
func New(conf *model.Conf) *fiber.App {
	app := fiber.New(&fiber.Settings{
		// Views:        pug.New("./public", ".pug"),
		ErrorHandler: handler.ErrorHandle(conf.Log),
	})

	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowHeaders:     []string{fiber.HeaderContentType, fiber.HeaderAuthorization},
		AllowMethods:     []string{fiber.MethodGet, fiber.MethodPost},
		AllowOrigins:     []string{"*"},
		MaxAge:           60 * 60 * 24,
		ExposeHeaders:    []string{fiber.HeaderContentType, fiber.HeaderAuthorization},
	}))
	app.Use(func(c *fiber.Ctx) {
		c.Locals("conf", conf)
		c.Next()
	})

	app.Get("/_/heath", func(c *fiber.Ctx) {
		c.SendStatus(fiber.StatusOK)
	})

	routing(app)

	return app
}
