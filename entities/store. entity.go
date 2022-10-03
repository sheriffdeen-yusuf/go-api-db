package models

import (
	"database/sql"
	"time"

)

type Store struct {
	Id           uint
	Name         string
	store        *string
	User  		User `gorm:"foreign:ID; constriant:onUpdate:CASCADE, onDelete:CASADE "`
	Password    string
	AccountType string
	Token      string
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
  }