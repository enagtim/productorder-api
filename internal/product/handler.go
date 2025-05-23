package product

import (
	"net/http"
	"order-api/pkg/messages"
	"order-api/pkg/req"
	"order-api/pkg/res"
	"strconv"
)

type ProductHandler struct {
	ProductService *ProductService
}

func NewProductHandler(router *http.ServeMux, service *ProductService) {
	handler := &ProductHandler{
		ProductService: service,
	}
	router.HandleFunc("POST /products", handler.CreateProduct())
	router.HandleFunc("GET /products", handler.GetAllProducts())
	router.HandleFunc("GET /products/{id}", handler.GetProductById())
	router.HandleFunc("PATCH /products/{id}", handler.UpdateProduct())
	router.HandleFunc("DELETE /products/{id}", handler.DeleteProduct())
}

func (h *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateDto](&w, r)
		if err != nil {
			return
		}
		product, err := h.ProductService.CreateProduct(*body)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, product, http.StatusCreated)
	}
}
func (h *ProductHandler) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil || offset < 0 {
			offset = 0
		}
		products, err := h.ProductService.GetAllProducts(limit, offset)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, products, http.StatusOK)
	}
}
func (h *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product, err := h.ProductService.GetProductById(uint(id))
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusNotFound)
			return
		}
		res.ResponseJson(w, product, http.StatusOK)
	}
}
func (h *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := req.HandleBody[ProductUpdateDto](&w, r)
		if err != nil {
			return
		}
		product, err := h.ProductService.UpdateProduct(uint(id), *body)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, product, http.StatusOK)
	}
}
func (h *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = h.ProductService.DeleteProduct(uint(id))
		if err != nil {
			messages.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, nil, http.StatusNoContent)
	}
}
