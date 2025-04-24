package repository

import (
	"finance-tracker/internal/model"
	"database/sql"
)

// TransactionRepository предоставляет методы для взаимодействия с базой данных.
type TransactionRepository struct {
	db *sql.DB
}

// NewTransactionRepository создаёт новый репозиторий.
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Add добавляет новую транзакцию.
func (r *TransactionRepository) Add(t model.Transaction) error {
	query := `INSERT INTO transactions (type, category, amount) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, t.Type, t.Category, t.Amount)
	return err
}

// List возвращает все транзакции.
func (r *TransactionRepository) List() ([]model.Transaction, error) {
	rows, err := r.db.Query(`SELECT id, type, category, amount, created_at FROM transactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(&t.ID, &t.Type, &t.Category, &t.Amount, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

// FilterByDate возвращает транзакции в заданном диапазоне дат.
func (r *TransactionRepository) FilterByDate(from, to string) ([]model.Transaction, error) {
	query := `SELECT id, type, category, amount, created_at FROM transactions WHERE DATE(created_at) BETWEEN ? AND ?`
	rows, err := r.db.Query(query, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(&t.ID, &t.Type, &t.Category, &t.Amount, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
