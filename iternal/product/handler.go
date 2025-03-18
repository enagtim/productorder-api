package product

import "net/http"

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, repo *ProductHandler) {
	handler := &ProductHandler{
		ProductRepository: repo.ProductRepository,
	}
	router.HandleFunc("POST /product/create", handler.Create())
	router.HandleFunc("GET /product/get", handler.Get())
	router.HandleFunc("PATCH /product/update", handler.Update())
	router.HandleFunc("DELETE /product/delete", handler.Delete())
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (h *ProductHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
