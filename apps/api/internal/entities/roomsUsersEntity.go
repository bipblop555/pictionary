package entities

import "time"

type RoomsUsersEntity struct {
	Id      string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	RoomsId string `json:"rooms_id"`
	UsersId string `json:"users_id"`
	Time    time.Time
}
