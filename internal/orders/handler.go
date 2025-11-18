package orders

import (
	"errors"
	"golang-ecom-api/internal/orders/json"
	"log/slog"
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

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var createOrderParams createOrderParams
	err := json.Read(r, &createOrderParams)
	if err != nil {
		slog.Error("failed to decode request body into createOrderParams", "error", err)
		http.Error(w, "failed to decode request body into createOrderParams", http.StatusBadRequest)
		return
	}

	order, err := h.svc.PlaceOrder(r.Context(), createOrderParams)
	if err != nil {
		slog.Error("failed to create order", "error", err)

		if errors.Is(err, errProductNotFind) {
			http.Error(w, "item does not exist", http.StatusBadRequest)
			return
		}

		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, order)
}
