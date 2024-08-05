package seeds

import (
	"log"
	"pictionary/internal/entities"
	"time"

	"gorm.io/gorm"
)

func SeedRoomsUsers(db *gorm.DB, roomsID []string, usersId []string) {
	roomsUsers := []entities.RoomsUsersEntity{
		{RoomsId: roomsID[0], UsersId: usersId[0], Time: time.Now()},
		{RoomsId: roomsID[0], UsersId: usersId[1], Time: time.Now()},
		{RoomsId: roomsID[0], UsersId: usersId[2], Time: time.Now()},
		{RoomsId: roomsID[0], UsersId: usersId[3], Time: time.Now()},
		{RoomsId: roomsID[0], UsersId: usersId[4], Time: time.Now()},
		{RoomsId: roomsID[0], UsersId: usersId[5], Time: time.Now()},
	}

	for _, roomUser := range roomsUsers {
		if err := db.Table("rooms_users").Create(&roomUser).Error; err != nil {
			log.Fatalf("Error inserting rooms_users %v", err)
		}
	}
}
