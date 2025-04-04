package user

type UserRequest struct {
	PhoneNumber string `json:"phonenumber" validate:"e164"`
}

type UserResponse struct {
	Token string `json:"token"`
}
