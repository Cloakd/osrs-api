package cmd

import (
	"encoding/json"
	"log"
	"strings"

	osrs_api "github.com/cloakd/osrs-api"
)

type (
	cacheItemScraper struct {
		cacheScraper
		Data []byte
	}

	cacheItem struct {
		Highalch      int         `json:"highalch"`
		Stackable     bool        `json:"stackable"`
		Lowalch       int         `json:"lowalch"`
		Cost          int         `json:"cost"`
		TradeableOnGe bool        `json:"tradeable_on_ge"`
		Equipable     bool        `json:"equipable"`
		Noteable      bool        `json:"noteable"`
		Noted         bool        `json:"noted"`
		Members       bool        `json:"members"`
		Name          string      `json:"name"`
		ID            int         `json:"id"`
		Placeholder   bool        `json:"placeholder"`
		LinkedID      interface{} `json:"linked_id"`
	}
)

const CACHE_ITEM_JSON = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/data/items-scraper.json"

func CacheItems() (osrs_api.ItemMap, error) {
	cms := cacheItemScraper{}
	var err error

	cms.Data, err = cms.Get(CACHE_ITEM_JSON)
	if err != nil {
		return nil, err
	}

	return cms.Parse()
}

func (cms *cacheItemScraper) Parse() (osrs_api.ItemMap, error) {
	items := map[string]cacheItem{}

	err := json.Unmarshal(cms.Data, &items)
	if err != nil {
		return nil, err
	}

	monMap := make(osrs_api.ItemMap, len(items))

	for _, mon := range items {
		//log.Printf("%v - %s (Level: %v)", id, mon.Name, mon.CombatLevel)
		monMap[mon.ID] = cms.cast(mon)
	}

	log.Printf("Cached %v Items", len(monMap))
	return monMap, nil
}

func (cms *cacheItemScraper) cast(mon cacheItem) *osrs_api.Item {
	return &osrs_api.Item{
		Id:      mon.ID,
		Name:    mon.Name,
		WikiUrl: strings.ReplaceAll(mon.Name, " ", "_"),
	}
}
