package seeds

import (
	"log"
	"pictionary/internal/entities"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) []string {
	var usersId []string

	users := []entities.UsersEntity{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
		{Name: "Louis"},
		{Name: "Karl"},
		{Name: "Peter"},
	}
	for _, user := range users {
		if err := db.Table("users").Create(&user).Error; err != nil {
			log.Fatalf("Error inserting users : %v", err)
		}
		usersId = append(usersId, user.Id)
	}

	return usersId
}
