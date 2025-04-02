package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

type JWTData struct {
	UserID uint
	Phone  string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) GenerateToken(phone string, userId uint) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userId,
		"phone":  phone,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) ParseToken(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone := t.Claims.(jwt.MapClaims)["phone"]
	userID := t.Claims.(jwt.MapClaims)["userID"]

	userIDFloat64, ok := userID.(float64)
	if !ok {
		return false, nil
	}
	userIDUint := uint(userIDFloat64)

	return t.Valid, &JWTData{
		UserID: userIDUint,
		Phone:  phone.(string),
	}
}
