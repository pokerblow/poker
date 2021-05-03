package poker

type Hand string

const (
	StraightFlush Hand = "Straight Flush"
	FourOfAKind Hand = "Four of a Kind"
	FullHouse Hand = "Full House"
	Flush Hand = "Flush"
	Straight Hand = "Straight"
	ThreeOfAKind Hand = "Three of a Kind"
	TwoPair Hand = "Two Pair"
	Pair Hand = "Pair"
	HighCard Hand = "High Card"
)
