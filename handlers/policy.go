package handlers

import (
	"pbac/db"
	"pbac/models"
	"pbac/pbac"

	"github.com/gofiber/fiber/v2"
)

func CreatePolicy(c *fiber.Ctx) error {
	var p models.Policy
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := db.DB.Exec(
		"INSERT INTO policies (user_id, resource_id, action_id, effect) VALUES (?,?,?,?)",
		p.UserID, p.ResourceID, p.ActionID, p.Effect,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := res.LastInsertId()
	p.ID = int(id)
	return c.JSON(p)
}

func CheckAccess(c *fiber.Ctx) error {
	var req struct {
		UserID     int `json:"user_id"`
		ResourceID int `json:"resource_id"`
		ActionID   int `json:"action_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	allowed, err := pbac.CheckAccess(req.UserID, req.ResourceID, req.ActionID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"allowed": allowed})
}

func ListPolicies(c *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT id, user_id, resource_id, action_id, effect FROM policies")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var policies []models.Policy
	for rows.Next() {
		var p models.Policy
		rows.Scan(&p.ID, &p.UserID, &p.ResourceID, &p.ActionID, &p.Effect)
		policies = append(policies, p)
	}
	return c.JSON(policies)
}
