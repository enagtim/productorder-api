package auth

type AuthorizationDto struct {
	Phone string `json:"phone" validate:"required,max=11"`
}
type TokenResponse struct {
	Token string `json:"token"`
}
