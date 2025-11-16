package products

import (
	"golang-ecom-api/internal/products/json"
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

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := struct {
		Products []string `json:"products"`
	}{
		Products: []string{"item-1", "item-2"},
	}

	json.Write(w, http.StatusOK, products)
}
