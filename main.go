package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `body:"body"`
}

func main() {
	fmt.Println("Hello")
	app := fiber.New()

	if envErr := godotenv.Load(".env"); envErr != nil {
		log.Fatal("Error on loading env")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"mssage": "Todo list", "todos": todos})
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

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		for i, todo := range todos {
			if todo.Id == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"messgae": "Todo deleted", "todo": todo})
			}
		}
		return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
	})

	app.Listen(":" + PORT)
}
