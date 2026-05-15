package orders

import (
	"net/http"

	jsonUtil "shop/internal/json"
)

type handler struct {
	service Service
}

// Constructor
func NewHandler(svc Service) *handler {
	return &handler{service: svc}
}

func (h *handler) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonUtil.Write(w, http.StatusOK, orders)
}

func (h *handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := jsonUtil.Read(r, &req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if errMsg := ValidateCreateOrderRequest(req); errMsg != "" {
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if err := h.service.CreateOrder(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
