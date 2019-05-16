package osrs_api

import "fmt"

type ItemDrop struct {
	Item        *Item
	Quantity    *Quantity
	Rarity      float32
	RarityNotes []*RarityNote
}

type Quantity struct {
	Min   int
	Max   int
	Noted bool
}

func (q *Quantity) String() string {
	str := fmt.Sprintf("%v", q.Max)

	if q.Min != q.Max {
		str = fmt.Sprintf("%v-%v", q.Min, q.Max)
	}

	if q.Noted {
		return fmt.Sprintf("%s (noted)", str)
	}

	return str
}
