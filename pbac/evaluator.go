package pbac

import (
	"pbac/db"
)

func CheckAccess(userID, resourceID, actionID int) (bool, error) {
	rows, err := db.DB.Query(`
        SELECT effect
        FROM policies
        WHERE user_id=? AND resource_id=? AND action_id=?`,
		userID, resourceID, actionID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	allowed := false
	for rows.Next() {
		var effect string
		if err := rows.Scan(&effect); err != nil {
			return false, err
		}

		if effect == "deny" {
			return false, nil // explicit deny wins
		} else if effect == "allow" {
			allowed = true
		}
	}

	return allowed, nil
}
