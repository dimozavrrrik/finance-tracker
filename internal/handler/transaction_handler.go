package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"finance-tracker/internal/model"
	"finance-tracker/internal/service"
)

// TransactionHandler обрабатывает HTTP-запросы.
type TransactionHandler struct {
	service *service.TransactionService
}

// NewTransactionHandler создаёт новый хендлер.
func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// HandleCreateTransaction — POST /transactions
func (h *TransactionHandler) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if err := h.service.AddTransaction(t); err != nil {
		http.Error(w, "Ошибка добавления транзакции", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// HandleGetTransactions — GET /transactions
func (h *TransactionHandler) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.ListTransactions()
	if err != nil {
		http.Error(w, "Ошибка получения транзакций", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(transactions)
}

// HandleFilterByDate — GET /transactions/filter?from=2024-01-01&to=2024-04-01
func (h *TransactionHandler) HandleFilterByDate(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if _, err := time.Parse("2006-01-02", from); err != nil {
		http.Error(w, "Неверная дата 'from'", http.StatusBadRequest)
		return
	}
	if _, err := time.Parse("2006-01-02", to); err != nil {
		http.Error(w, "Неверная дата 'to'", http.StatusBadRequest)
		return
	}

	transactions, err := h.service.FilterByDate(from, to)
	if err != nil {
		http.Error(w, "Ошибка фильтрации", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(transactions)
}
