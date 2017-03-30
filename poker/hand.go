package poker

import "time"

type Hand struct {
	ID             int64
	PlayerID       int
	PlayerName     string
	Cards          []*Card
	GameNumber     string
	BetAmount      int64
	NetAmount      int64
	ValidBetAmount int64
	PayoutAmount   int64
	Currency       int
	IsWinner       bool
	PlayType       int
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}

func New() *Hand {
	return nil
}

func Cards() []*Card {
	return nil
}

func AddCards([]*Card) error {
	return nil
}
