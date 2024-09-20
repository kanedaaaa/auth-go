package types

type SignupPyaload struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
