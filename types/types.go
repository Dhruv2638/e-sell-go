package types

type RegisterUserPayLoad struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"emial"`
	Password  string `json:"password"`
}
