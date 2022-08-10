package models

import "time"

type Page struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Body      string `json:"body"`
	IdPage    string `json:"id_page"`
}
