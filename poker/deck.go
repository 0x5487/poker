package poker

type Deck struct {
	Cards []*Card
}

func (deck *Deck) Pop() *Card {
	return nil
}

func (dec *Deck) PopMulti(n int) []*Card {
	return nil
} 