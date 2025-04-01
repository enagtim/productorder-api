package order

type OrderService struct {
	OrderRepository *OrderRepository
}

func NewOrderService(orderRepository *OrderRepository) *OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
	}
}

func (s *OrderService) CreateOrder(userID uint, productIDs []uint) (*Order, error) {
	order, err := s.OrderRepository.CreateOrder(userID, productIDs)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) FindOrderByID(orderID, userID uint) (*Order, error) {
	order, err := s.OrderRepository.FindOrderByID(orderID, userID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetAllProductsByUser(userID uint) (*[]Order, error) {
	orders, err := s.OrderRepository.GetAllProductsByUser(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
