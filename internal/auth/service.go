package auth

import (
	"errors"
	"order-api/internal/user"
	"order-api/pkg/session"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) CreateUser(phone string) (*user.User, error) {
	existedUser, _ := s.UserRepository.FindByPhone(phone)
	if existedUser != nil {
		return existedUser, nil
	}
	sessionId, err := session.GenerateSessionId()
	if err != nil {
		return nil, errors.New("error generation sessionId")
	}
	user := &user.User{
		Phone:     phone,
		SessionId: sessionId,
	}
	_, err = s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *AuthService) VerifyUser(phone, sessionId string) (string, error) {
	user, err := s.UserRepository.FindByPhone(phone)
	if err != nil {
		return "", errors.New("user not found")
	}
	if user.SessionId != sessionId {
		return "", errors.New("error user sessionId")
	}
	return user.SessionId, nil
}
