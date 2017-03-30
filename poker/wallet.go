package poker

import "time"

type Wallet struct {
	ID        int
	PlayerID  int
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
