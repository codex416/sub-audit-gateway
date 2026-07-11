package database


import (
	"time"
)



type AuditLog struct {


	ID uint `json:"id"`



	Token string `json:"token"`



	IP string `json:"ip"`



	Country string `json:"country"`



	UserAgent string `json:"user_agent"`



	Client string `json:"client"`



	Path string `json:"path"`



	Status int `json:"status"`



	CreatedAt time.Time `json:"created_at"`

}
