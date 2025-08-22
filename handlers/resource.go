package handlers

import (
	"pbac/db"
	"pbac/models"

	"github.com/gofiber/fiber/v2"
)

func CreateResource(c *fiber.Ctx) error {
	var rsc models.Resource
	if err := c.BodyParser(&rsc); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := db.DB.Exec("INSERT INTO resources (name) VALUES (?)", rsc.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := res.LastInsertId()
	rsc.ID = int(id)
	return c.JSON(rsc)
}

func ListResources(c *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT id, name FROM resources")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var resources []models.Resource
	for rows.Next() {
		var r models.Resource
		rows.Scan(&r.ID, &r.Name)
		resources = append(resources, r)
	}
	return c.JSON(resources)
}
