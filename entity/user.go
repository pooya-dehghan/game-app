package entity

type User struct {
	ID             uint   `json:"id"`
	PhoneNumber    string `json:"phone_number"`
	Avatar         string `json:"avatar"`
	Name           string `json:"name"`
	HashedPassword string `json:"-"`
}
