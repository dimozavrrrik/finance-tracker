# Finance Tracker
Простое приложение для учета финансовых транзакций с REST API на Go

## Возможности:
- Добавление транзакций
- Фильтрация по датам
- Хранение в  SQLite
- RESTful API

### Требования
- Go 1.21+
- SQLite3

```bash
git clone https://github.com/dimozavrrrik/finance-tracker.git
cd finance-tracker
go run cmd/main.go