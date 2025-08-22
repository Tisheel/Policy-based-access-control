package handlers

import (
	"pbac/db"
	"pbac/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := db.DB.Exec("INSERT INTO users (username) VALUES (?)", u.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := res.LastInsertId()
	u.ID = int(id)
	return c.JSON(u)
}

func ListUsers(c *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT id, username FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Username)
		users = append(users, u)
	}
	return c.JSON(users)
}
