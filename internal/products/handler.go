package products

import (
	"golang-ecom-api/internal/products/json"
	"log"
	"net/http"
)

type handler struct {
	svc Service
}

func NewHandler(svc Service) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.Write(w, http.StatusOK, products)
}
