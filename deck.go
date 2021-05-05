package poker

import (
	"math/rand"
	"time"
)

var fullDeck *Deck

func init() {
	fullDeck = &Deck{initializeFullCards()}
	rand.Seed(time.Now().UnixNano())
}

type Deck struct {
	cards []card
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.Shuffle()
	return deck
}

func (deck *Deck) Shuffle() {
	deck.cards = make([]card, len(fullDeck.cards))
	copy(deck.cards, fullDeck.cards)
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

func (deck *Deck) Draw(n int) []card {
	cards := make([]card, n)
	copy(cards, deck.cards[:n])
	deck.cards = deck.cards[n:]
	return cards
}

func (deck *Deck) Empty() bool {
	return len(deck.cards) == 0
}

func initializeFullCards() []card {
	var cards []card

	for _, rank := range strRanks {
		for suit := range charSuitToIntSuit {
			cards = append(cards, cardFromStr(string(rank)+string(suit)))
		}
	}

	return cards
}
