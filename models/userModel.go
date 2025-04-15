package models

type User struct {
	ID         int    `json:"id"`
	ProfilePic string `json:"profile_pic"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
