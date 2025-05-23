package product

type ProductCreateDto struct {
	Name        string   `json:"name" validate:"required,min=1,max=50"`
	Description string   `json:"description" validate:"required,min=1,max=200"`
	Images      []string `json:"images,omitempty" validate:"dive,url"`
	Price       float64  `json:"price" validate:"required,gte=0"`
	Discount    float64  `json:"discount" validate:"gte=0,lte=100"`
}

type ProductUpdateDto struct {
	Name        string   `json:"name" validate:"max=50"`
	Description string   `json:"description" validate:"max=200"`
	Images      []string `json:"images" validate:"dive,url"`
	Price       float64  `json:"price" validate:"gte=0"`
	Discount    float64  `json:"discount" validate:"gte=0,lte=100"`
}
