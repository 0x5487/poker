package domain

import (
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Suit  string
	Value int
}

type Deck []Card

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func NewDeck() *Deck {
	suits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	deck := Deck{}

	for _, suit := range suits {
		for _, value := range values {
			card := Card{
				Suit:  suit,
				Value: value,
			}
			deck = append(deck, card)
		}
	}

	return &deck
}

type Hand struct {
	Cards []*Card
}

func (h *Hand) AddCard(card *Card) {
	h.Cards = append(h.Cards, card)
}

type Player struct {
	Name string
	Hand Hand
	Bet  int
}

type Board struct {
	Flop   Hand
	Turn   *Card
	River  *Card
	Pot    int
	Dealer int
}

func determineWinner(players []*Player, board Board) *Player {
	var bestHandRank int
	var winningPlayer *Player

	for _, player := range players {
		// Evaluate player's hand using the board cards
		playerHand := append(player.Hand.Cards, board.Flop.Cards...)
		playerHand = append(playerHand, board.Turn, board.River)
		handRank := evaluateHand(playerHand)

		// Update winningPlayer and bestHandRank if current player's hand is better
		if handRank > bestHandRank {
			bestHandRank = handRank
			winningPlayer = player
		}
	}

	return winningPlayer
}

func evaluateHand(cards []*Card) int {
	// Check for straight flush
	if rank := checkStraightFlush(cards); rank != 0 {
		return rank
	}

	// Check for four of a kind
	if rank := checkFourOfAKind(cards); rank != 0 {
		return rank
	}

	// Check for full house
	if rank := checkFullHouse(cards); rank != 0 {
		return rank
	}

	// Check for flush
	if rank := checkFlush(cards); rank != 0 {
		return rank
	}

	// Check for straight
	if rank := checkStraight(cards); rank != 0 {
		return rank
	}

	// Check for three of a kind
	if rank := checkThreeOfAKind(cards); rank != 0 {
		return rank
	}

	// Check for two pairs
	if rank := checkTwoPairs(cards); rank != 0 {
		return rank
	}

	// Check for pair
	if rank := checkPair(cards); rank != 0 {
		return rank
	}

	// If no hand ranking was found, return the rank of the highest card
	return checkHighCard(cards)
}

func checkStraightFlush(cards []*Card) int {
	// Create a map to count the number of cards of each suit
	counts := make(map[string]int)
	for _, card := range cards {
		counts[card.Suit]++
	}

	// Check if there is a suit with at least 5 cards
	var flushSuit string
	for suit, count := range counts {
		if count >= 5 {
			flushSuit = suit
			break
		}
	}
	if flushSuit == "" {
		return 0
	}

	// Create a slice of cards of the flush suit
	var flushCards []*Card
	for _, card := range cards {
		if card.Suit == flushSuit {
			flushCards = append(flushCards, card)
		}
	}

	// Sort the cards by value
	sort.Slice(flushCards, func(i, j int) bool {
		return flushCards[i].Value < flushCards[j].Value
	})

	// Check if there is a straight within the flush suit
	for i := 0; i <= len(flushCards)-5; i++ {
		if flushCards[i+4].Value-flushCards[i].Value == 4 {
			return 9 + flushCards[i+4].Value // Return rank of straight flush, with highest card as tiebreaker
		}
	}

	return 0
}

func checkFourOfAKind(hand []*Card) int {
	// Sort hand by rank
	sort.Slice(hand, func(i, j int) bool {
		return hand[i].Value < hand[j].Value
	})

	// Check if there is a four of a kind
	for i := 0; i <= len(hand)-4; i++ {
		if hand[i].Value == hand[i+1].Value && hand[i+1].Value == hand[i+2].Value && hand[i+2].Value == hand[i+3].Value {
			return hand[i].Value
		}
	}

	return 0
}

func checkFullHouse(hand []*Card) int {
	// First check if there are three cards of the same rank
	threeOfAKind := checkThreeOfAKind(hand)
	if threeOfAKind == 0 {
		return 0
	}

	// Then check if there is a pair with a different rank
	pair := checkPair(getCardsByRank(hand, threeOfAKind))
	if pair == 0 {
		return 0
	}

	// Return the rank of the three of a kind plus the rank of the pair
	return threeOfAKind + pair
}

func getCardsByRank(hand []*Card, rank int) []*Card {
	var cards []*Card
	for _, card := range hand {
		if card.Value == rank {
			cards = append(cards, card)
		}
	}
	return cards
}

func checkFlush(hand []*Card) int {
	flushSuit := hand[0].Suit
	for _, card := range hand[1:] {
		if card.Suit != flushSuit {
			return 0
		}
	}
	return 6
}

func checkStraight(hand []*Card) int {
	ranks := make([]int, len(hand))
	for i, card := range hand {
		ranks[i] = card.Value
	}
	sort.Ints(ranks)

	// Check if there is an Ace at the bottom and King at the top of the straight
	if ranks[0] == 1 && ranks[1] == 10 && ranks[2] == 11 && ranks[3] == 12 && ranks[4] == 13 {
		return 5 // Ace-high straight
	}

	// Check for other straights
	for i := 0; i < len(ranks)-1; i++ {
		if ranks[i]+1 != ranks[i+1] {
			return 0 // Not a straight
		}
	}

	return 4 // Straight
}

func checkHighCard(hand []*Card) int {
	highCardRank := 0
	for _, card := range hand {
		cardRank := card.Value
		if cardRank > highCardRank {
			highCardRank = cardRank
		}
	}
	return highCardRank
}

func checkPair(cards []*Card) int {
	rankMap := make(map[int]int)
	for _, card := range cards {
		rankMap[card.Value]++
	}
	for rank, count := range rankMap {
		if count == 2 {
			return rank
		}
	}
	return 0
}

func checkTwoPairs(cards []*Card) int {
	pairCount := 0
	highCard := 0

	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[i].Value == cards[j].Value {
				pairCount++
				if cards[i].Value > highCard {
					highCard = cards[i].Value
				}
			}
		}
	}

	if pairCount == 2 {
		return highCard
	}

	return 0
}

func checkThreeOfAKind(cards []*Card) int {
	cardCounts := make(map[int]int)

	for _, card := range cards {
		cardCounts[card.Value]++
	}

	for cardVal, count := range cardCounts {
		if count == 3 {
			return cardVal
		}
	}

	return 0
}
