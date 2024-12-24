package libs

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CleanText(text string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9@.]+`)
	return regex.ReplaceAllString(text, "")
}

func ExtractToken(c *fiber.Ctx) string {
	bearer := c.Get("Authorization")
	token := strings.Replace(bearer, "Bearer ", "", 1)
	return token

}
