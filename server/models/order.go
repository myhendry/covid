package models

type Order struct {
	ID uint
	Name string `json:"name"`
	Remarks string `json:"remarks"`
	Valid bool `json:"valid" gorm:"default:true"`
	Total float64 `json:"total" gorm:"-"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID uint
	OrderID uint `json:"order-id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

func (order *Order) GetTotal() float64 {
	var total float64 = 0

	for _, item := range order.OrderItems {
		total += item.Price
	}

	return total
}