package auth

import (
	"net/http"
	"order-api/configs"
	"order-api/pkg/jwt"
	"order-api/pkg/messages"
	"order-api/pkg/req"
	"order-api/pkg/res"
)

type AuthHandler struct {
	AuthService *AuthService
	Config      *configs.Config
}
type AuthHandlerDeps struct {
	AuthService *AuthService
	Config      *configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
		Config:      deps.Config,
	}
	router.HandleFunc("POST /auth", handler.AuthUser())
}

func (h *AuthHandler) AuthUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[AuthorizationDto](&w, r)
		if err != nil {
			return
		}
		createdUser, err := h.AuthService.CreateUser(body.Phone)

		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionId, err := h.AuthService.VerifyUser(createdUser.Phone, createdUser.SessionId)

		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusNotFound)
			return
		}

		token, err := jwt.NewJWT(h.Config.Auth.Secret).GenerateToken(sessionId)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TokenResponse{
			Token: token,
		}
		res.ResponseJson(w, data, http.StatusOK)

	}
}
