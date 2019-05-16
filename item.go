package osrs_api

type (
	ItemMap map[int]*Item

	Item struct {
		Id      int
		Name    string
		WikiUrl string

		Highalch      int         `json:"highalch"`
		Stackable     bool        `json:"stackable"`
		Lowalch       int         `json:"lowalch"`
		Cost          int         `json:"cost"`
		TradeableOnGe bool        `json:"tradeable_on_ge"`
		Equipable     bool        `json:"equipable"`
		Noteable      bool        `json:"noteable"`
		Noted         bool        `json:"noted"`
		Members       bool        `json:"members"`
		ID            int         `json:"id"`
		Placeholder   bool        `json:"placeholder"`
		LinkedID      interface{} `json:"linked_id"`
	}
)
