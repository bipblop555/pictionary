package database

import (
	"clean/internal/database/seeds"
	"clean/internal/env"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	env.Init()

	conninfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("dbport"), os.Getenv("dbuser"), os.Getenv("dbpassword"), os.Getenv("dbname"))
	conn, err := gorm.Open(postgres.Open(conninfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to db ", err)
	}

	var categoryIds = seeds.SeedCategory(conn)
	seeds.SeedProduct(conn, categoryIds)

	return conn
}
