package poker

import "time"

type Seat struct {
	ID     int
	Player Player
}

type Table struct {
	Code      string
	Dealer    *Dealer
	Seats     []*Seat
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

func (tbl *Table) GameInfo() {

}

func (tbl *Table) Stand(seatID int) {

}

func (tbl *Table) Sit(seatID int, player *Player) {

}

func (tbl *Table) Next() {

}
