package handler

import (
	"boilerplate/internal/model"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func ErrorHandle(log *logrus.Entry) func(*fiber.Ctx, error) {
	return func(c *fiber.Ctx, err error) {
		e := err.(*model.Error)
		log.WithFields(logrus.Fields{
			"status-code": e.Code,
			"debug":       e.Debug,
		}).Error(e.Message)

		c.Status(e.Code).JSON(fiber.Map{
			"message": e.Message,
		})
	}
}
