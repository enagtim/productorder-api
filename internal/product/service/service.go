package service

import (
	"order-api/internal/product/dto"
	"order-api/internal/product/model"
	"order-api/internal/product/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: repo}
}

func (s *ProductService) CreateProduct(req dto.ProductCreateDto) (*model.Product, error) {
	product := model.NewProduct(req.Name, req.Description, req.Images, req.Price, req.Discount)
	return s.ProductRepository.Create(product)
}
func (s *ProductService) GetAllProducts(limit, offset int) (*[]model.Product, error) {
	return s.ProductRepository.GetProducts(limit, offset)
}
func (s *ProductService) GetProductById(id uint) (*model.Product, error) {
	return s.ProductRepository.GetById(id)
}
func (s *ProductService) UpdateProduct(id uint, req dto.ProductUpdateDto) (*model.Product, error) {
	_, err := s.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	product := &model.Product{
		Model:       gorm.Model{ID: id},
		Name:        req.Name,
		Description: req.Description,
		Images:      req.Images,
		Price:       req.Price,
		Discount:    req.Discount,
	}
	return s.ProductRepository.Update(product)
}
func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.ProductRepository.GetById(id)
	if err != nil {
		return err
	}
	return s.ProductRepository.Delete(id)
}
