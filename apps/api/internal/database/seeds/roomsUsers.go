package seeds

import (
	"log"
	"pictionary/internal/entities"

	"gorm.io/gorm"
)

func SeedRoomsUsers(db *gorm.DB, roomsID []string, usersId []string) {
	roomsUsers := []entities.RoomsUsersEntity{
		{RoomsId: roomsID[0], UsersId: usersId[0]},
		{RoomsId: roomsID[0], UsersId: usersId[1]},
		{RoomsId: roomsID[0], UsersId: usersId[2]},
		{RoomsId: roomsID[0], UsersId: usersId[3]},
		{RoomsId: roomsID[0], UsersId: usersId[4]},
		{RoomsId: roomsID[0], UsersId: usersId[5]},
	}

	for _, roomUser := range roomsUsers {
		if err := db.Table("rooms_users").Create(&roomUser).Error; err != nil {
			log.Fatalf("Error inserting rooms_users %v", err)
		}
	}
}
