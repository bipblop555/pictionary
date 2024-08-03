package entities

type RoomsEntity struct {
	Id string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
}
