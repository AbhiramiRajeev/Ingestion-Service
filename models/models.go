package models

type Event struct {
	EventType string   `json:"event_type" binding:"required"`
	UserID    string   `json:"user_id" binding:"required"`
	IpAddress string   `json:"ip_address" binding:"required"`
	Status	  string   `json:"status" binding:"required"`
	Timestamp string   `json:"timestamp" binding:"required"`
} 
