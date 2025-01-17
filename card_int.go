package poker

import "github.com/pokerblow/poker/cardb"

type card int32

var (
	intRanks [13]int32
	strRanks = "23456789TJQKA"
	primes   = []int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
)

var (
	charRankToIntRank = map[uint8]int32{}
	charSuitToIntSuit = map[uint8]int32{
		's': 1, // spades
		'h': 2, // hearts
		'd': 4, // diamonds
		'c': 8, // clubs
	}
	intSuitToCharSuit = "xshxdxxxc"
)

var (
	prettySuits = map[int]string{
		1: "\u2660", // spades
		2: "\u2764", // hearts
		4: "\u2666", // diamonds
		8: "\u2663", // clubs
	}
	prettyReds = [...]int{2, 4}
)

func init() {
	for i := 0; i < 13; i++ {
		intRanks[i] = int32(i)
	}

	for i := range strRanks {
		charRankToIntRank[strRanks[i]] = intRanks[i]
	}
}

func cardFromStr(s string) card {
	rankInt := charRankToIntRank[s[0]]
	suitInt := charSuitToIntSuit[s[1]]
	rankPrime := primes[rankInt]

	bitRank := int32(1) << uint32(rankInt) << 16
	suit := suitInt << 12
	rank := rankInt << 8

	return card(bitRank | suit | rank | rankPrime)
}

func cardFromB(c cardb.Card) card {
	return cardFromStr(c.String())
}

func toCards(cards []card) []cardb.Card {
	result := make([]cardb.Card, 0, len(cards))
	for _, c := range cards {
		result = append(result, cardb.Card(c.String()))
	}
	return result
}

func (c card) String() string {
	return string(strRanks[c.Rank()]) + string(intSuitToCharSuit[c.Suit()])
}

func (c card) Rank() int32 {
	return (int32(c) >> 8) & 0xF
}

func (c card) Suit() int32 {
	return (int32(c) >> 12) & 0xF
}

func (c card) BitRank() int32 {
	return (int32(c) >> 16) & 0x1FFF
}

func (c card) Prime() int32 {
	return int32(c) & 0x3F
}

func primeProductFromHand(cards []card) int32 {
	product := int32(1)

	for _, card := range cards {
		product *= (int32(card) & 0xFF)
	}

	return product
}

func primeProductFromRankBits(rankBits int32) int32 {
	product := int32(1)

	for _, i := range intRanks {
		if rankBits&(1<<uint(i)) != 0 {
			product *= primes[i]
		}
	}

	return product
}
