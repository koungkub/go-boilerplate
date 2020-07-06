package handler

import (
	"boilerplate/internal/model"
	"boilerplate/internal/worker"
	"context"

	"github.com/gofiber/fiber"
)

func GetReview() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		conf := c.Locals("conf").(*model.Conf)

		w := worker.NewReview(conf.DB)
		w.Get(context.TODO(), "1")
		if 1 == 1 {
			c.Next(model.NewError(fiber.StatusUnprocessableEntity, "error", "debug naja"))
			return
		}

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	}
}
