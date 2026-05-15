package orders

type CreateOrderRequest struct {
	CustomerID int32       `json:"customerId"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductID int32 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}
