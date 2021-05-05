package poker

import (
	"fmt"
	"github.com/pokerblow/poker/cardb"
)

type Result struct {
	Rank int32
	Hand
	// Best five cards from evaluation.
	Cards []cardb.Card
}

var table *lookupTable

func init() {
	table = newLookupTable()
}

func rankClass(rank int32) int32 {
	targets := [...]int32{
		maxStraightFlush,
		maxFourOfAKind,
		maxFullHouse,
		maxFlush,
		maxStraight,
		maxThreeOfAKind,
		maxTwoPair,
		maxPair,
		maxHighCard,
	}

	if rank < 0 {
		panic(fmt.Sprintf("rank %d is less than zero", rank))
	}

	for _, target := range targets {
		if rank <= target {
			return maxToRankClass[target]
		}
	}

	panic(fmt.Sprintf("rank %d is unknown", rank))
}

func rankHand(rank int32) Hand {
	return rankClassToString[rankClass(rank)]
}

func rankString(rank int32) string {
	return string(rankClassToString[rankClass(rank)])
}

func EvalCards(cards []cardb.Card) Result {
	a := make([]card, 0, len(cards))
	for _, card := range cards {
		a = append(a, cardFromB(card))
	}
	return eval(a)
}

func Eval(cards []string) Result {
	a := make([]card, 0, len(cards))
	for _, card := range cards {
		a = append(a, cardFromStr(card))
	}
	return eval(a)
}

func eval(incoming []card) Result {
	var bestCards []card
	var bestRank int32
	switch len(incoming) {
	case 5:
		bestRank, bestCards = five(incoming...), incoming
	case 6:
		bestRank, bestCards = six(incoming...)
	case 7:
		bestRank, bestCards = seven(incoming...)
	default:
		panic("Only support 5, 6 and 7 cards.")
	}
	return Result{
		Rank: bestRank,
		Hand: rankHand(bestRank),
		Cards: toCards(bestCards),
	}
}

func evaluate(cards []card) int32 {
	return eval(cards).Rank
}

func five(cards ...card) int32 {
	if cards[0]&cards[1]&cards[2]&cards[3]&cards[4]&0xF000 != 0 {
		handOR := (cards[0] | cards[1] | cards[2] | cards[3] | cards[4]) >> 16
		prime := primeProductFromRankBits(int32(handOR))
		return table.flushLookup[prime]
	}

	prime := primeProductFromHand(cards)
	return table.unsuitedLookup[prime]
}

func six(cards ...card) (int32, []card) {
	var bestCombRank int32 = maxHighCard
	var bestComb []card
	targets := make([]card, len(cards))

	for i := 0; i < len(cards); i++ {
		copy(targets, cards)
		targets := append(targets[:i], targets[i+1:]...)

		score := five(targets...)
		if score < bestCombRank {
			bestCombRank = score
			bestComb = targets
		}
	}

	return bestCombRank, bestComb
}

func seven(cards ...card) (int32, []card) {
	var bestCombRank int32 = maxHighCard
	var bestComb []card
	targets := make([]card, len(cards))

	for i := 0; i < len(cards); i++ {
		copy(targets, cards)
		targets := append(targets[:i], targets[i+1:]...)

		score, bestCards := six(targets...)
		if score < bestCombRank {
			bestCombRank = score
			bestComb = bestCards
		}
	}

	return bestCombRank, bestComb
}
