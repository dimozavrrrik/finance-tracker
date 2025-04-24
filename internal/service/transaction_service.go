package service

import (
	"finance-tracker/internal/model"
	"finance-tracker/internal/repository"
)

// TransactionService описывает бизнес-логику приложения.
type TransactionService struct {
	repo *repository.TransactionRepository
}

// NewTransactionService создаёт новый сервис.
func NewTransactionService(r *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

// AddTransaction добавляет транзакцию через репозиторий.
func (s *TransactionService) AddTransaction(t model.Transaction) error {
	return s.repo.Add(t)
}

// ListTransactions возвращает все транзакции.
func (s *TransactionService) ListTransactions() ([]model.Transaction, error) {
	return s.repo.List()
}

// FilterByDate фильтрует транзакции по диапазону дат.
func (s *TransactionService) FilterByDate(from, to string) ([]model.Transaction, error) {
	return s.repo.FilterByDate(from, to)
}
