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
	err := h.svc.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{
		Products: []string{"item-1", "item-2"},
	}

	json.Write(w, http.StatusOK, products)
}
