package daifugo

import (
	"errors"
	"sort"

	"github.com/fuuki/board/resource"
)

// Suit is a suit of a playing card.
type Suit int

const (
	SuitSpade Suit = iota
	SuitHeart
	SuitDiamond
	SuitClub
)

// Card is a playing card.
type Card struct {
	CardID string `json:"id"`
	Rank   int    `json:"rank"`
	Suit   Suit   `json:"suit"`
}

var _ resource.Card = (*Card)(nil)

// ID returns a card ID.
func (c *Card) ID() resource.CardID {
	return resource.CardID(c.CardID)
}

// LineType is a type of a set of cards.
type LineType int

const (
	LineTypeInvalid LineType = iota
	LineTypeSameRank
	LineTypeStairs
)

// LinePower is a power of a set of cards.
type LinePower struct {
	LineType LineType
	Count    int
	Power    int
}

// SortedCards is a sorted list of cards.
type SortedCards struct {
	list []*Card
}

// NewSortedCards returns a new sorted list of cards.
func NewSortedCards(cards []*Card) (*SortedCards, error) {
	if len(cards) == 0 {
		return nil, errors.New("no cards")
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank == cards[j].Rank {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Rank < cards[j].Rank
	})
	return &SortedCards{
		list: cards,
	}, nil
}

// GetLinePower judges the power of a set of cards.
// If the set of cards is invalid, it returns nil.
func (s *SortedCards) Power() *LinePower {
	t := s.getLineType()
	if t == LineTypeInvalid {
		return nil
	}
	return &LinePower{
		LineType: t,
		Count:    len(s.list),
		Power:    s.list[0].Rank,
	}
}

// GetLineType judges the type of a set of cards.
func (s *SortedCards) getLineType() LineType {
	isSameRank := func(cards []*Card) bool {
		rank := cards[0].Rank
		for _, c := range cards[1:] {
			if c.Rank != rank {
				return false
			}
		}
		return true
	}
	isStairs := func(cards []*Card) bool {
		// 3枚以上の階段のみ許可する
		if len(cards) < 3 {
			return false
		}
		rank := cards[0].Rank
		suit := cards[0].Suit
		for _, c := range cards[1:] {
			rank++
			if c.Rank != rank || c.Suit != suit {
				return false
			}
		}
		return true
	}

	if isSameRank(s.list) {
		return LineTypeSameRank
	}
	if isStairs(s.list) {
		return LineTypeStairs
	}
	return LineTypeInvalid
}

var shortDeck = []*Card{
	{CardID: "s1", Rank: 1, Suit: SuitSpade},
	{CardID: "s2", Rank: 2, Suit: SuitSpade},
	{CardID: "s3", Rank: 3, Suit: SuitSpade},
	{CardID: "s4", Rank: 4, Suit: SuitSpade},
	{CardID: "s5", Rank: 5, Suit: SuitSpade},
	{CardID: "h1", Rank: 1, Suit: SuitHeart},
	{CardID: "h2", Rank: 2, Suit: SuitHeart},
	{CardID: "h3", Rank: 3, Suit: SuitHeart},
	{CardID: "h4", Rank: 4, Suit: SuitHeart},
	{CardID: "h5", Rank: 5, Suit: SuitHeart},
}

var deck = []*Card{
	{CardID: "s1", Rank: 1, Suit: SuitSpade},
	{CardID: "s2", Rank: 2, Suit: SuitSpade},
	{CardID: "s3", Rank: 3, Suit: SuitSpade},
	{CardID: "s4", Rank: 4, Suit: SuitSpade},
	{CardID: "s5", Rank: 5, Suit: SuitSpade},
	{CardID: "s6", Rank: 6, Suit: SuitSpade},
	{CardID: "s7", Rank: 7, Suit: SuitSpade},
	{CardID: "s8", Rank: 8, Suit: SuitSpade},
	{CardID: "s9", Rank: 9, Suit: SuitSpade},
	{CardID: "s10", Rank: 10, Suit: SuitSpade},
	{CardID: "s11", Rank: 11, Suit: SuitSpade},
	{CardID: "s12", Rank: 12, Suit: SuitSpade},
	{CardID: "s13", Rank: 13, Suit: SuitSpade},
	{CardID: "h1", Rank: 1, Suit: SuitHeart},
	{CardID: "h2", Rank: 2, Suit: SuitHeart},
	{CardID: "h3", Rank: 3, Suit: SuitHeart},
	{CardID: "h4", Rank: 4, Suit: SuitHeart},
	{CardID: "h5", Rank: 5, Suit: SuitHeart},
	{CardID: "h6", Rank: 6, Suit: SuitHeart},
	{CardID: "h7", Rank: 7, Suit: SuitHeart},
	{CardID: "h8", Rank: 8, Suit: SuitHeart},
	{CardID: "h9", Rank: 9, Suit: SuitHeart},
	{CardID: "h10", Rank: 10, Suit: SuitHeart},
	{CardID: "h11", Rank: 11, Suit: SuitHeart},
	{CardID: "h12", Rank: 12, Suit: SuitHeart},
	{CardID: "h13", Rank: 13, Suit: SuitHeart},
	{CardID: "d1", Rank: 1, Suit: SuitDiamond},
	{CardID: "d2", Rank: 2, Suit: SuitDiamond},
	{CardID: "d3", Rank: 3, Suit: SuitDiamond},
	{CardID: "d4", Rank: 4, Suit: SuitDiamond},
	{CardID: "d5", Rank: 5, Suit: SuitDiamond},
	{CardID: "d6", Rank: 6, Suit: SuitDiamond},
	{CardID: "d7", Rank: 7, Suit: SuitDiamond},
	{CardID: "d8", Rank: 8, Suit: SuitDiamond},
	{CardID: "d9", Rank: 9, Suit: SuitDiamond},
	{CardID: "d10", Rank: 10, Suit: SuitDiamond},
	{CardID: "d11", Rank: 11, Suit: SuitDiamond},
	{CardID: "d12", Rank: 12, Suit: SuitDiamond},
	{CardID: "d13", Rank: 13, Suit: SuitDiamond},
	{CardID: "c1", Rank: 1, Suit: SuitClub},
	{CardID: "c2", Rank: 2, Suit: SuitClub},
	{CardID: "c3", Rank: 3, Suit: SuitClub},
	{CardID: "c4", Rank: 4, Suit: SuitClub},
	{CardID: "c5", Rank: 5, Suit: SuitClub},
	{CardID: "c6", Rank: 6, Suit: SuitClub},
	{CardID: "c7", Rank: 7, Suit: SuitClub},
	{CardID: "c8", Rank: 8, Suit: SuitClub},
	{CardID: "c9", Rank: 9, Suit: SuitClub},
	{CardID: "c10", Rank: 10, Suit: SuitClub},
	{CardID: "c11", Rank: 11, Suit: SuitClub},
	{CardID: "c12", Rank: 12, Suit: SuitClub},
	{CardID: "c13", Rank: 13, Suit: SuitClub},
}
