package poker

import "time"

type Game struct {
	ID        int64
	Number    string
	BetAmount int64
	Currency  int
	WinnerID  int
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
