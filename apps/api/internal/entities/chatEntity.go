package entities

type Content struct {
	Content string `json:"content,omitempty"`
}

type ChatEntity struct {
	Content Content
	UsersId string `json:"users_id"`
}
