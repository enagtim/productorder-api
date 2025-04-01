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
}

func NewOrderHandler(router *http.ServeMux, deps *OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderService: deps.OrderService,
	}
	router.Handle("POST /orders", middleware.IsAuthed(handler.CreateOrder(), deps.Config))
	router.Handle("GET /orders/{id}", middleware.IsAuthed(handler.FindOrderByID(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuthed(handler.GetAllProductsByUser(), deps.Config))
}
func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(middleware.ContextUserID).(uint)
		if !ok {
			messages.SendJSONError(w, "User not found with context", http.StatusUnauthorized)
			return
		}
		body, err := req.HandleBody[OrderCreateDto](&w, r)
		if err != nil {
			return
		}
		order, err := h.OrderService.CreateOrder(userId, body.ProductsIDs)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, order, http.StatusCreated)
	}
}
func (h *OrderHandler) FindOrderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(middleware.ContextUserID).(uint)
		if !ok {
			messages.SendJSONError(w, "User not found with context", http.StatusUnauthorized)
			return
		}
		orderIdString := r.PathValue("id")

		orderID, err := strconv.Atoi(orderIdString)

		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		order, err := h.OrderService.FindOrderByID(uint(orderID), userId)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, order, http.StatusOK)
	}
}
func (h *OrderHandler) GetAllProductsByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(middleware.ContextUserID).(uint)
		if !ok {
			messages.SendJSONError(w, "User not found with context", http.StatusUnauthorized)
			return
		}
		orders, err := h.OrderService.GetAllProductsByUser(userId)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, orders, http.StatusOK)
	}
}
