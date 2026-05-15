package orders

import "fmt"

// ValidateCreateOrderRequest validates the CreateOrderRequest and returns an error message if invalid, or "" if valid.
func ValidateCreateOrderRequest(req CreateOrderRequest) string {
	if req.CustomerID <= 0 {
		return "customerId is required and must be positive"
	}

	if len(req.Items) == 0 {
		return "items is required and must contain at least one item"
	}

	for i, item := range req.Items {
		if item.ProductID <= 0 {
			return fmt.Sprintf("items[%d].productId must be positive", i)
		}
		if item.Quantity <= 0 {
			return fmt.Sprintf("items[%d].quantity must be positive", i)
		}
	}

	return ""
}
