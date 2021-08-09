package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slaveofcode/okegani/v1/users"
)

func BuildRoutes(router fiber.Router) {
	u := router.Group("/users")
	u.Post("/register", users.Register)
}
