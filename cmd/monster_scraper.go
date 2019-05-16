package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	osrs_api "github.com/cloakd/osrs-api"
	"github.com/cloakd/osrs-api/wiki"
)

const (
	DROP_TAG = "DropsLine/Sandbox"
)

type (
	Scraper struct {
		Client http.Client
	}

	MonsterScraper struct {
		Scraper *Scraper
	}

	MonsterIndex []osrs_api.Monster
)

func NewScraper() *Scraper {
	return &Scraper{
		Client: http.Client{},
	}
}

func NewMonsterScraper() *MonsterScraper {
	return &MonsterScraper{
		Scraper: NewScraper(),
	}
}

func (s *MonsterScraper) Scrape() {
	log.Printf("Starting Monster Scrape")

	index := wiki.GetMonsters()
	log.Printf("Monster Urls Retrieved: %v", len(index))

	for _, m := range index {

		body, err := s.Scraper.source(m)
		if err != nil {
			log.Printf("Unable to get source (%s) - %s ", m, err)
			continue
		}

		p := wiki.NewParser(body)
		results := p.Tags(DROP_TAG)

		log.Printf("Obtained %v drop results for %s", len(results), m)
		for _, r := range results {
			log.Printf("Drop Result: %s", r)

			//TODO Parse result
		}
	}

}

func (s *Scraper) source(id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(wiki.WIKI_EDIT, id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	str, err := ioutil.ReadAll(resp.Body)

	return str, nil
}

func (s *MonsterScraper) parseDrop(result string) *osrs_api.ItemDrop {
	return &osrs_api.ItemDrop{
		//Item        *Item
		//Quantity    *Quantity
		//Rarity      float32
		//RarityNotes []*RarityNote
	}
}
