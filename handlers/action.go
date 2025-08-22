package handlers

import (
	"pbac/db"
	"pbac/models"

	"github.com/gofiber/fiber/v2"
)

func CreateAction(c *fiber.Ctx) error {
	var act models.Action
	if err := c.BodyParser(&act); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := db.DB.Exec("INSERT INTO actions (name) VALUES (?)", act.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := res.LastInsertId()
	act.ID = int(id)
	return c.JSON(act)
}

func ListActions(c *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT id, name FROM actions")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var actions []models.Action
	for rows.Next() {
		var a models.Action
		rows.Scan(&a.ID, &a.Name)
		actions = append(actions, a)
	}
	return c.JSON(actions)
}
