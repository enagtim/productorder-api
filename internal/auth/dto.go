package auth

type AuthorizationDto struct {
	Phone string `json:"phone" validate:"required,min=11,max=11"`
}
type TokenResponse struct {
	Token string `json:"token"`
}
