package entities

import "time"

type RoomsEntity struct {
	Id   string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Time time.Time
}
