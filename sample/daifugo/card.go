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
	id   string
	Rank int
	Suit Suit
}

var _ resource.Card = (*Card)(nil)

// ID returns a card ID.
func (c *Card) ID() resource.CardID {
	return resource.CardID(c.id)
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
	{id: "s1", Rank: 1, Suit: SuitSpade},
	{id: "s2", Rank: 2, Suit: SuitSpade},
	{id: "s3", Rank: 3, Suit: SuitSpade},
	{id: "s4", Rank: 4, Suit: SuitSpade},
	{id: "s5", Rank: 5, Suit: SuitSpade},
	{id: "h1", Rank: 1, Suit: SuitHeart},
	{id: "h2", Rank: 2, Suit: SuitHeart},
	{id: "h3", Rank: 3, Suit: SuitHeart},
	{id: "h4", Rank: 4, Suit: SuitHeart},
	{id: "h5", Rank: 5, Suit: SuitHeart},
}

var deck = []*Card{
	{id: "s1", Rank: 1, Suit: SuitSpade},
	{id: "s2", Rank: 2, Suit: SuitSpade},
	{id: "s3", Rank: 3, Suit: SuitSpade},
	{id: "s4", Rank: 4, Suit: SuitSpade},
	{id: "s5", Rank: 5, Suit: SuitSpade},
	{id: "s6", Rank: 6, Suit: SuitSpade},
	{id: "s7", Rank: 7, Suit: SuitSpade},
	{id: "s8", Rank: 8, Suit: SuitSpade},
	{id: "s9", Rank: 9, Suit: SuitSpade},
	{id: "s10", Rank: 10, Suit: SuitSpade},
	{id: "s11", Rank: 11, Suit: SuitSpade},
	{id: "s12", Rank: 12, Suit: SuitSpade},
	{id: "s13", Rank: 13, Suit: SuitSpade},
	{id: "h1", Rank: 1, Suit: SuitHeart},
	{id: "h2", Rank: 2, Suit: SuitHeart},
	{id: "h3", Rank: 3, Suit: SuitHeart},
	{id: "h4", Rank: 4, Suit: SuitHeart},
	{id: "h5", Rank: 5, Suit: SuitHeart},
	{id: "h6", Rank: 6, Suit: SuitHeart},
	{id: "h7", Rank: 7, Suit: SuitHeart},
	{id: "h8", Rank: 8, Suit: SuitHeart},
	{id: "h9", Rank: 9, Suit: SuitHeart},
	{id: "h10", Rank: 10, Suit: SuitHeart},
	{id: "h11", Rank: 11, Suit: SuitHeart},
	{id: "h12", Rank: 12, Suit: SuitHeart},
	{id: "h13", Rank: 13, Suit: SuitHeart},
	{id: "d1", Rank: 1, Suit: SuitDiamond},
	{id: "d2", Rank: 2, Suit: SuitDiamond},
	{id: "d3", Rank: 3, Suit: SuitDiamond},
	{id: "d4", Rank: 4, Suit: SuitDiamond},
	{id: "d5", Rank: 5, Suit: SuitDiamond},
	{id: "d6", Rank: 6, Suit: SuitDiamond},
	{id: "d7", Rank: 7, Suit: SuitDiamond},
	{id: "d8", Rank: 8, Suit: SuitDiamond},
	{id: "d9", Rank: 9, Suit: SuitDiamond},
	{id: "d10", Rank: 10, Suit: SuitDiamond},
	{id: "d11", Rank: 11, Suit: SuitDiamond},
	{id: "d12", Rank: 12, Suit: SuitDiamond},
	{id: "d13", Rank: 13, Suit: SuitDiamond},
	{id: "c1", Rank: 1, Suit: SuitClub},
	{id: "c2", Rank: 2, Suit: SuitClub},
	{id: "c3", Rank: 3, Suit: SuitClub},
	{id: "c4", Rank: 4, Suit: SuitClub},
	{id: "c5", Rank: 5, Suit: SuitClub},
	{id: "c6", Rank: 6, Suit: SuitClub},
	{id: "c7", Rank: 7, Suit: SuitClub},
	{id: "c8", Rank: 8, Suit: SuitClub},
	{id: "c9", Rank: 9, Suit: SuitClub},
	{id: "c10", Rank: 10, Suit: SuitClub},
	{id: "c11", Rank: 11, Suit: SuitClub},
	{id: "c12", Rank: 12, Suit: SuitClub},
	{id: "c13", Rank: 13, Suit: SuitClub},
}
