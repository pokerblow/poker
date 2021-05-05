package cardb

import "fmt"

type Card string

func New(s string) Card {
	if len(s) != 2 {
		panic(fmt.Sprintf("invalid card string %s", s))
	}
	// todo add pattern validation
	return Card(s)
}

func NewFrom(f Face, s Suit) Card {
	return Card(string(f) + string(s))
}

func (c Card) Face() Face {
	return Face(c[0])
}

// Deuce = 0, Three = 1, ..., Ace = 12
func (c Card) FaceRank() int {
	return faceRanks[c.Face()]
}

func (c Card) Suit() Suit {
	return Suit(c[1])
}

func (c Card) String() string {
	return string(c)
}
