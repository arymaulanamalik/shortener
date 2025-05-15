package url

import "time"

type URL struct {
	ID         string    `json:"id" db:"id"`
	ShortCode  string    `json:"short_code" db:"short_code"`
	Original   string    `json:"original" db:"original"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ClickCount int       `json:"click_count" db:"click_count"`
}
