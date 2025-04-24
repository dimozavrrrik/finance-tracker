package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"finance-tracker/internal/handler"
	"finance-tracker/internal/repository"
	"finance-tracker/internal/service"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Ошибка при подключении к БД:", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT NOT NULL,
		category TEXT NOT NULL,
		amount REAL NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatal("Ошибка при создании таблицы:", err)
	}

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	h := handler.NewTransactionHandler(svc)

	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.HandleCreateTransaction(w, r)
		case http.MethodGet:
			h.HandleGetTransactions(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/transactions/filter", h.HandleFilterByDate)
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Добро пожаловать в Finance Tracker!"))
	//})
	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
