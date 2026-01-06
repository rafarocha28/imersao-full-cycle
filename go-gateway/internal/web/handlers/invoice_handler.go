package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")

	var input dto.CreateInvoiceInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = apiKey

	output, err := h.invoiceService.Create(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *InvoiceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invoice ID is required", http.StatusBadRequest)
		return
	}

	output, err := h.invoiceService.GetById(id, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *InvoiceHandler) List(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")

	output, err := h.invoiceService.ListByAccountApiKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}