package seeds

import (
	"log"
	"pictionary/internal/entities"

	"gorm.io/gorm"
)

func SeedRooms(db *gorm.DB) []string {
	var roomsId []string
	rooms := []entities.RoomsEntity{
		{Id: ""},
		{Id: ""},
	}
	for _, room := range rooms {
		if err := db.Table("rooms").Create(&room).Error; err != nil {
			log.Fatalf("Error inserting rooms : %v", err)
		}
		roomsId = append(roomsId, room.Id)
	}
	return roomsId
}
