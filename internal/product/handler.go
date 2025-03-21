package product

import (
	"net/http"
	"order-api/pkg/messages"
	"order-api/pkg/req"
	"order-api/pkg/res"
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
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		product := NewProduct(body.Name, body.Description, body.Images, body.Price, body.Discount)
		createdProduct, err := h.ProductRepository.Create(product)
		if err != nil {
			messages.SendJSONEError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.ResponseJson(w, createdProduct, http.StatusCreated)
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
		res.ResponseJson(w, product, http.StatusOK)
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
		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		product, err := h.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
			Price:       body.Price,
			Discount:    body.Discount,
		})
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
		res.ResponseJson(w, nil, http.StatusNoContent)
	}
}
