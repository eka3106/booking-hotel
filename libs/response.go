package libs

import "github.com/gofiber/fiber/v2"

func ResponseError(c *fiber.Ctx, message interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"errors": message})

}

func ResponseSuccess(c *fiber.Ctx, message interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"data": message})
}
