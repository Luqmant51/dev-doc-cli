package models

type User struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	UserID      string `json:"user_id"`
}
