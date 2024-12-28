package libs

import "github.com/gofiber/fiber/v2"

func ResponseError(c *fiber.Ctx, message interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"errors": message})

}

func ResponseSuccess(c *fiber.Ctx, message interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"data": message})
}

func ResponsePagination(c *fiber.Ctx, data interface{}, totalData int64, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"total_data": totalData, "data": data })
}
