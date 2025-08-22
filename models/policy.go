package models

type Policy struct {
	ID         int    `json:"id" db:"id"`
	UserID     int    `json:"user_id" db:"user_id"`
	ResourceID int    `json:"resource_id" db:"resource_id"`
	ActionID   int    `json:"action_id" db:"action_id"`
	Effect     string `json:"effect" db:"effect"`
}
