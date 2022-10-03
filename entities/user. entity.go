package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name        string  `gorm:"type:varchar(255)" json:"name"`
	Email       string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	PhoneNumber string
	Password    string  `gorm:"->;<-;not null" json:"-"`
	AccountType string `gorm:"-" json:"token,omitempty"`
	Token       string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time

	
}
