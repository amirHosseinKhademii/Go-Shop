package products

import (
	"log"
	"net/http"
	"shop/internal/json"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	parsedId, err := strconv.ParseInt(param, 16, 32)
	var id int32 = int32(parsedId)
	products, err := h.service.GetProductById(r.Context(), id)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)
}
