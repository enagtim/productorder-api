package product

import (
	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepository *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{ProductRepository: repo}
}

func (s *ProductService) CreateProduct(req ProductCreateDto) (*Product, error) {
	product := NewProduct(req.Name, req.Description, req.Images, req.Price, req.Discount)
	return s.ProductRepository.Create(product)
}
func (s *ProductService) GetAllProducts(limit, offset int) (*[]Product, error) {
	return s.ProductRepository.GetProducts(limit, offset)
}
func (s *ProductService) GetProductById(id uint) (*Product, error) {
	return s.ProductRepository.GetById(id)
}
func (s *ProductService) UpdateProduct(id uint, req ProductUpdateDto) (*Product, error) {
	_, err := s.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	product := &Product{
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
