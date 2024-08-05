package seeds

import (
	"log"
	"pictionary/internal/entities"
	"time"

	"gorm.io/gorm"
)

func SeedRooms(db *gorm.DB) []string {
	var roomsId []string
	rooms := []entities.RoomsEntity{
		{Id: "", Time: time.Now()},
		{Id: "", Time: time.Now()},
	}
	for _, room := range rooms {
		if err := db.Table("rooms").Create(&room).Error; err != nil {
			log.Fatalf("Error inserting rooms : %v", err)
		}
		roomsId = append(roomsId, room.Id)
	}
	return roomsId
}
