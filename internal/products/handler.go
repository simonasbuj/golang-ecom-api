package products

import (
	"errors"
	"golang-ecom-api/internal/products/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var errProductIdNotProvided = errors.New("product id not provided in url")

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

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")
	if productIDStr == "" {
		log.Println(errProductIdNotProvided)
		http.Error(w, errProductIdNotProvided.Error(), http.StatusBadRequest)
		return
	}

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		log.Println("invalid product ID:", err)
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.svc.GetProductByID(r.Context(), productID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
