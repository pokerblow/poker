package poker

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankIntegers(t *testing.T) {
	for i := 0; i < 13; i++ {
		assert.Equal(t, int32(i), intRanks[i])
	}
}

func TestNewCard(t *testing.T) {
	assert.Equal(t, card(268446761), cardFromStr("Ah"))
	assert.Equal(t, card(134224677), cardFromStr("Ks"))
}

func TestMarshalJSON(t *testing.T) {
	cards := []card{
		cardFromStr("Ah"),
		cardFromStr("Kh"),
		cardFromStr("Qh"),
		cardFromStr("Jh"),
		cardFromStr("Th"),
	}

	b, err := json.Marshal(cards)
	assert.NoError(t, err)
	assert.Equal(t, `["Ah","Kh","Qh","Jh","Th"]`, string(b))
}

func TestUnmarshalJSON(t *testing.T) {
	var cards []card
	data := `["Ah","Kh","Qh","Jh","Th"]`

	err := json.Unmarshal([]byte(data), &cards)
	assert.NoError(t, err)
	assert.Contains(t, cards, cardFromStr("Ah"))
	assert.Contains(t, cards, cardFromStr("Kh"))
	assert.Contains(t, cards, cardFromStr("Qh"))
	assert.Contains(t, cards, cardFromStr("Jh"))
	assert.Contains(t, cards, cardFromStr("Th"))
}

func TestString(t *testing.T) {
	assert.Equal(t, "3s", cardFromStr("3s").String())
}

func TestBitRank(t *testing.T) {
	assert.Equal(t, int32(2048), cardFromStr("Ks").BitRank())
}

func TestPrimeProductFromHand(t *testing.T) {
	assert.Equal(t, int32(20387), primeProductFromHand([]card{8398611, 134236965, 33564957}))
	assert.Equal(t, int32(61161), primeProductFromHand([]card{8398611, 134236965, 33564957, 135427}))
	assert.Equal(t, int32(183483), primeProductFromHand([]card{8398611, 134236965, 33564957, 135427, 139523}))
}

func TestPrimeProductFromRankBits(t *testing.T) {
	assert.Equal(t, int32(2), primeProductFromRankBits(1))
	assert.Equal(t, int32(42), primeProductFromRankBits(11))
	assert.Equal(t, int32(110), primeProductFromRankBits(21))
	assert.Equal(t, int32(2310), primeProductFromRankBits(31))
	assert.Equal(t, int32(4290), primeProductFromRankBits(55))
	assert.Equal(t, int32(1785), primeProductFromRankBits(78))
	assert.Equal(t, int32(1326), primeProductFromRankBits(99))
	assert.Equal(t, int32(34034), primeProductFromRankBits(121))
	assert.Equal(t, int32(30107), primeProductFromRankBits(344))
}
