package poker

import "time"

type Table struct {
	Code      string
	Dealer    *Dealer
	Players   []*Player
	maxPlayer int
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

func (tbl *Table) GameInfo() {

}

func (tbl *Table) Stand(player *Player) {

}

func (tbl *Table) Sit(player *Player) {

}

func (tbl *Table) Next() {

}
