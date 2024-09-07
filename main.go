package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `body:"body"`
}

func main() {
	fmt.Println("Hello")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"mssage": "hello world"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"message": "Todo body is required"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(fiber.Map{"todos": todos})
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i := 0; i < len(todos); i++ {
			iid, _ := strconv.Atoi(id)
			if iid == todos[i].Id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(404).JSON(fiber.Map{"message": "Todo udpated", "todo": todos[i]})
			}
		}
		return c.Status(400).JSON(fiber.Map{"messgae": "Todo not found"})
	})

	app.Listen(":4000")
}
