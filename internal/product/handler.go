package product

import (
	"net/http"
	"order-api/internal/product/dto"
	"order-api/internal/product/service"
	"order-api/pkg/messages"
	"order-api/pkg/req"
	"order-api/pkg/res"
	"strconv"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(router *http.ServeMux, service *service.ProductService) {
	handler := &ProductHandler{
		ProductService: service,
	}
	router.HandleFunc("POST /product/create", handler.CreateProduct())
	router.HandleFunc("GET /products", handler.GetAllProducts())
	router.HandleFunc("GET /product/get/{id}", handler.GetProductById())
	router.HandleFunc("PATCH /product/update/{id}", handler.UpdateProduct())
	router.HandleFunc("DELETE /product/delete/{id}", handler.DeleteProduct())
}

func (h *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[dto.ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		product, err := h.ProductService.CreateProduct(*body)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
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
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
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
			messages.SendJSONEError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product, err := h.ProductService.GetProductById(uint(id))
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusNotFound)
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
			messages.SendJSONEError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := req.HandleBody[dto.ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		product, err := h.ProductService.UpdateProduct(uint(id), *body)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
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
			messages.SendJSONEError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = h.ProductService.DeleteProduct(uint(id))
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, nil, http.StatusNoContent)
	}
}
