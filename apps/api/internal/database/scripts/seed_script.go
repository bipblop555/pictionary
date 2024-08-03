package main

import (
	"fmt"
	"log"
	"os"
	"pictionary/internal/database/seeds"
	"pictionary/internal/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	env.Init()

	conninfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("dbport"), os.Getenv("dbuser"), os.Getenv("dbpassword"), os.Getenv("dbname"))
	conn, err := gorm.Open(postgres.Open(conninfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to db ", err)
	}

	var roomsIds = seeds.SeedRooms(conn)
	var usersIds = seeds.SeedUsers(conn)

	seeds.SeedRoomsUsers(conn, roomsIds, usersIds)

}
