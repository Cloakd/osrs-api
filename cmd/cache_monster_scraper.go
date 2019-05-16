package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	osrs_api "github.com/cloakd/osrs-api"
)

const CACHE_JSON = "https://raw.githubusercontent.com/osrsbox/osrsbox-db/master/data/attackable-npcs.json"

type (
	cacheMonsterScraper struct {
		cacheScraper
		Data []byte
	}

	cacheScraper struct {
	}

	cacheMonster struct {
		ID                     int           `json:"id"`
		Rotation               int           `json:"rotation"`
		Name                   string        `json:"name"`
		Models                 []int32       `json:"models"`
		StanceAnimation        int32         `json:"stanceAnimation"`
		AnInt2165              int           `json:"anInt2165"`
		TileSpacesOccupied     int8          `json:"tileSpacesOccupied"`
		WalkAnimation          int32         `json:"walkAnimation"`
		Rotate90RightAnimation int32         `json:"rotate90RightAnimation"`
		ABool2170              bool          `json:"aBool2170"`
		ResizeX                int16         `json:"resizeX"`
		Contrast               int           `json:"contrast"`
		Rotate180Animation     int32         `json:"rotate180Animation"`
		VarbitIndex            int           `json:"varbitIndex"`
		Options                []interface{} `json:"options"`
		RenderOnMinimap        bool          `json:"renderOnMinimap"`
		CombatLevel            int16         `json:"combatLevel"`
		Rotate90LeftAnimation  int32         `json:"rotate90LeftAnimation"`
		ResizeY                int16         `json:"resizeY"`
		HasRenderPriority      bool          `json:"hasRenderPriority"`
		Ambient                int           `json:"ambient"`
		HeadIcon               int           `json:"headIcon"`
		VarpIndex              int           `json:"varpIndex"`
		IsClickable            bool          `json:"isClickable"`
		AnInt2189              int           `json:"anInt2189"`
		ABool2190              bool          `json:"aBool2190"`
	}
)

func CacheMonsters() (osrs_api.MonsterMap, error) {
	cms := cacheMonsterScraper{}
	var err error

	cms.Data, err = cms.Get(CACHE_JSON)
	if err != nil {
		return nil, err
	}

	return cms.Parse()
}

func (cs *cacheScraper) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func (cms *cacheMonsterScraper) Parse() (osrs_api.MonsterMap, error) {
	monsters := map[string]cacheMonster{}

	err := json.Unmarshal(cms.Data, &monsters)
	if err != nil {
		return nil, err
	}

	monMap := make(osrs_api.MonsterMap, len(monsters))

	for _, mon := range monsters {
		//log.Printf("%v - %s (Level: %v)", id, mon.Name, mon.CombatLevel)
		monMap[mon.ID] = cms.cast(mon)
	}

	log.Printf("Cached %v Monsters", len(monMap))
	return monMap, nil
}

func (cms *cacheMonsterScraper) cast(mon cacheMonster) *osrs_api.Monster {
	return &osrs_api.Monster{
		Id:          mon.ID,
		Name:        mon.Name,
		WikiUrl:     strings.ReplaceAll(mon.Name, " ", "_"),
		CombatLevel: mon.CombatLevel,
		Models:      mon.Models,

		Animations: osrs_api.Animation{
			Stance:        mon.StanceAnimation,
			Walk:          mon.WalkAnimation,
			Rotate90Left:  mon.Rotate90LeftAnimation,
			Rotate90Right: mon.Rotate90RightAnimation,
			Rotate180:     mon.Rotate180Animation,
		},

		Dimensions: osrs_api.Dimension{
			TileSpace: mon.TileSpacesOccupied,
			ResizeX:   mon.ResizeX,
			ResizeY:   mon.ResizeY,
		},
	}
}
