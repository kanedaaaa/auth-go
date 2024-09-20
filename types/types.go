package types

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type SignupPyaload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
