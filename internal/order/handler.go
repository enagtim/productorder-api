package order

import (
	"net/http"
	"order-api/configs"
	"order-api/pkg/messages"
	"order-api/pkg/middleware"
	"order-api/pkg/req"
	"order-api/pkg/res"
	"strconv"
)

type OrderHandlerDeps struct {
	OrderService *OrderService
	Config       *configs.Config
}
type OrderHandler struct {
	OrderService *OrderService
	Config       *configs.Config
}

func NewProductHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderService: deps.OrderService,
		Config:       deps.Config,
	}
	router.Handle("POST /orders", middleware.IsAuthed(handler.CreateOrder(), deps.Config))
}
func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userQueryId := r.URL.Query().Get("userId")
		userId, err := strconv.Atoi(userQueryId)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := req.HandleBody[OrderCreateDto](&w, r)
		if err != nil {
			return
		}
		order, err := h.OrderService.CreateOrder(uint(userId), body.ProductsIDs)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, order, http.StatusCreated)
	}
}
