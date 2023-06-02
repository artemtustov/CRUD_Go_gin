package domain

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Email       string `json:"Email"`
	PasswordEnc string `json:"PasswordEnc"`
}

type UserRq struct {
	Email       string `json:"Email"`
	PasswordEnc string `json:"PasswordEnc"`
}
