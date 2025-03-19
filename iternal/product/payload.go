package product

type ProductPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Discount    uint   `json:"discount"`
}
