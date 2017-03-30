package poker

import "time"

type Dealer struct {
	ID        int
	Name      string
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
