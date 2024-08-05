package seeds

import (
	"log"
	"pictionary/internal/entities"
	"time"

	"gorm.io/gorm"
)

func SeedChat(db *gorm.DB, usersId []string) {

	chats := []entities.ChatEntity{
		{Content: entities.Content{Content: "Coucou les gars"}, UsersId: usersId[0], Time: time.Now()},
		{Content: entities.Content{Content: "Coucou les gars"}, UsersId: usersId[1], Time: time.Now()},
		{Content: entities.Content{Content: "Coucou les gars"}, UsersId: usersId[2], Time: time.Now()},
		{Content: entities.Content{Content: "Coucou les gars"}, UsersId: usersId[3], Time: time.Now()},
		{Content: entities.Content{Content: "Coucou les gars"}, UsersId: usersId[0], Time: time.Now()},
	}

	for _, chat := range chats {
		if err := db.Table("chat").Create(&chat).Error; err != nil {
			log.Fatalf("Error sending chat %v", err)
		}
	}
}
