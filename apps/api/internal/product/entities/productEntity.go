package entities

import (
	"time"
)

type ProductEntity struct {
	Id           string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CategoriesId string    `json:"categories_id" gorm:"categories_id"`
	Name         string    `json:"name"`
	Price        float32   `json:"price"`
	Quantity     int       `json:"quantity"`
	Created_at   time.Time `json:"created_at"`
	Deleted_at   time.Time `json:"deleted_at"`
}
