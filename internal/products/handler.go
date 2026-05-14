package products

import (
	"log"
	"net/http"
	"shop/internal/json"
)

type handler struct {
	service Service
}

// Constructor
func NewHandler(svc Service) *handler {
	return &handler{
		service: svc,
	}
}

func (h *handler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)

}
