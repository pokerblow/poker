package poker

import (
	"fmt"
)

type Result struct {
	Rank int32
	Hand
}

var table *lookupTable

func init() {
	table = newLookupTable()
}

func RankClass(rank int32) int32 {
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
	return rankClassToString[RankClass(rank)]
}

//Deprecated, Result has the Hand.
func RankString(rank int32) string {
	return string(rankClassToString[RankClass(rank)])
}

func Eval(cards []string) Result {
	var a []Card
	for _, card := range cards {
		a = append(a, NewCard(card))
	}
	rank := evalRank(a)
	return Result{
		Rank: rank,
		Hand: rankHand(rank),
	}
}

func evalRank(cards []Card) int32{
	switch len(cards) {
	case 5:
		return five(cards...)
	case 6:
		return six(cards...)
	case 7:
		return seven(cards...)
	default:
		panic("Only support 5, 6 and 7 cards.")
	}
}

//Deprecated, use Eval.
func Evaluate(cards []Card) int32 {
	return evalRank(cards)
}

func five(cards ...Card) int32 {
	if cards[0]&cards[1]&cards[2]&cards[3]&cards[4]&0xF000 != 0 {
		handOR := (cards[0] | cards[1] | cards[2] | cards[3] | cards[4]) >> 16
		prime := primeProductFromRankBits(int32(handOR))
		return table.flushLookup[prime]
	}

	prime := primeProductFromHand(cards)
	return table.unsuitedLookup[prime]
}

func six(cards ...Card) int32 {
	var minimum int32 = maxHighCard
	targets := make([]Card, len(cards))

	for i := 0; i < len(cards); i++ {
		copy(targets, cards)
		targets := append(targets[:i], targets[i+1:]...)

		score := five(targets...)
		if score < minimum {
			minimum = score
		}
	}

	return minimum
}

func seven(cards ...Card) int32 {
	var minimum int32 = maxHighCard
	targets := make([]Card, len(cards))

	for i := 0; i < len(cards); i++ {
		copy(targets, cards)
		targets := append(targets[:i], targets[i+1:]...)

		score := six(targets...)
		if score < minimum {
			minimum = score
		}
	}

	return minimum
}
