package product

import (
	"encoding/json"
	"net/http"
	"order-api/pkg/messages"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, repo *ProductHandler) {
	handler := &ProductHandler{
		ProductRepository: repo.ProductRepository,
	}
	router.HandleFunc("POST /product/create", handler.CreateProduct())
	router.HandleFunc("GET /product/get/{id}", handler.GetProductById())
	router.HandleFunc("PATCH /product/update/{id}", handler.UpdateProduct())
	router.HandleFunc("DELETE /product/delete/{id}", handler.DeleteProduct())
}

func (h *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload ProductPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		product := NewProduct(payload.Name, payload.Description, payload.Price, payload.Discount)
		createdProduct, err := h.ProductRepository.Create(product)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdProduct)
	}
}
func (h *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := h.ProductRepository.GetById(uint(id))
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}
func (h *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.ProductRepository.GetById(uint(id))
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusNotFound)
			return
		}
		var payload ProductPayload
		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := h.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        payload.Name,
			Description: payload.Description,
			Price:       payload.Price,
			Discount:    payload.Discount,
		})
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}
func (h *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.ProductRepository.GetById(uint(id))
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.ProductRepository.Delete(uint(id))
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
