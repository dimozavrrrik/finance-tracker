package model

// Transaction представляет одну финансовую операцию (расход или доход).
type Transaction struct {
	ID        int     `json:"id"`
	Type      string  `json:"type"`     // тип: зачисление или списание
	Category  string  `json:"category"` // категория, например "еда", "зарплата"
	Amount    float64 `json:"amount"`   // сумма
	CreatedAt string  `json:"created_at"`
}
