package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type RequestData struct {
	JariJariLingkaran int `json:"jari-jari-lingkaran"`
}

type ResponseData struct {
	LuasLingkaran     float64 `json:"luas-lingkaran"`
	KelilingLingkaran float64 `json:"keliling-lingkaran"`
}

func main() {
	app := fiber.New()

	const port = ":3000"

	app.Post("/", func(c *fiber.Ctx) error {
		request := new(RequestData)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		if request.JariJariLingkaran <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid radius value",
			})
		}

		response := new(ResponseData)

		response.LuasLingkaran = 3.14 * float64(request.JariJariLingkaran*request.JariJariLingkaran)
		response.KelilingLingkaran = 2 * 3.14 * float64(request.JariJariLingkaran)

		return c.JSON(response)
	})

	app.Listen(port)
}
