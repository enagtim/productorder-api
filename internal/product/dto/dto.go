package dto

type ProductCreateRequest struct {
	Name        string   `json:"name" validate:"required,max=50"`
	Description string   `json:"description" validate:"required,max=200"`
	Images      []string `json:"images,omitempty" validate:"dive,url"`
	Price       float64  `json:"price" validate:"required,gte=0"`
	Discount    float64  `json:"discount" validate:"gte=0,lte=100"`
}

type ProductUpdateRequest struct {
	Name        string   `json:"name" validate:"max=50"`
	Description string   `json:"description" validate:"max=200"`
	Images      []string `json:"images" validate:"dive,url"`
	Price       float64  `json:"price" validate:"gte=0"`
	Discount    float64  `json:"discount" validate:"gte=0,lte=100"`
}
