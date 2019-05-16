package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	osrs_api "github.com/cloakd/osrs-api"
	"github.com/cloakd/osrs-api/wiki"
)

const (
	DROP_TAG = "DropsLine"
)

type (
	Scraper struct {
		Wait   sync.WaitGroup
		Client http.Client
	}

	MonsterScraper struct {
		MonsterMap *osrs_api.MonsterMap
		Scraper    *Scraper
	}

	MonsterIndex []osrs_api.Monster
)

func NewScraper() *Scraper {
	return &Scraper{
		Client: http.Client{},
	}
}

func NewMonsterScraper(monsterMap *osrs_api.MonsterMap) *MonsterScraper {
	return &MonsterScraper{
		MonsterMap: monsterMap,
		Scraper:    NewScraper(),
	}
}

func (s *MonsterScraper) Scrape(done chan *osrs_api.Monster) {
	log.Printf("Starting Monster Scrape")

	jobs := make(chan *osrs_api.Monster, len(*s.MonsterMap))

	for w := 1; w <= 3; w++ {
		go s.worker(w, jobs, done)
	}

	s.Scraper.Wait = sync.WaitGroup{}
	for _, m := range *s.MonsterMap {

		s.Scraper.Wait.Add(1)
		jobs <- m

		//body, err := s.Scraper.source(m.WikiUrl)
		//if err != nil {
		//	log.Printf("Unable to get source (%s) - %s ", m, err)
		//	continue
		//}
		//
		//p := wiki.NewParser(body)
		//results := p.Tags(DROP_TAG)
		//
		//log.Printf("Obtained %v drop results for %s", len(results), m.Name)
		//for _, r := range results {
		//	log.Printf("Drop Result: %s", r)
		//
		//	drop := s.parseDrop(r)
		//	m.ItemDrops.Add(drop)
		//}
		//
		//done <- m
	}
	close(jobs)

	s.Scraper.Wait.Wait()
}

func (s *MonsterScraper) worker(id int, jobs <-chan *osrs_api.Monster, done chan<- *osrs_api.Monster) {
	for m := range jobs {
		body, err := s.Scraper.source(m.WikiUrl)
		if err != nil {
			log.Printf("Unable to get source (%s) - %s ", m, err)
			continue
		}

		p := wiki.NewParser(body)
		results := p.Tags(DROP_TAG)

		log.Printf("Obtained %v drop results for %s", len(results), m.Name)
		for _, r := range results {
			log.Printf("Drop Result: %s", r)

			drop := s.parseDrop(r)
			m.ItemDrops.Add(drop)
		}

		s.Scraper.Wait.Done()
		done <- m
	}
}

func (s *Scraper) source(id string) ([]byte, error) {
	log.Printf("Querying Url: %s", fmt.Sprintf(wiki.WIKI_EDIT, id))
	resp, err := http.Get(fmt.Sprintf(wiki.WIKI_EDIT, id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	str, err := ioutil.ReadAll(resp.Body)

	return str, nil
}

func (s *MonsterScraper) parseDrop(result []byte) *osrs_api.ItemDrop {

	r := strings.ReplaceAll(string(result), "{{DropsLine/Sandbox|", "")
	r = strings.ReplaceAll(r, "{{DropsLine/sandbox|", "")
	r = strings.ReplaceAll(r, "{{DropsLine|", "")
	r = strings.ReplaceAll(r, "{{(m)}}", "")
	r = strings.ReplaceAll(r, "}}", "")
	lines := strings.Split(r, "|")

	d := osrs_api.ItemDrop{}

	for _, l := range lines {
		kv := strings.Split(l, "=")
		if len(kv) != 2 {
			log.Printf("WARN: Markdown Error: %s - %s", l, kv)
			continue
		}

		switch kv[0] {
		case "Name":
			d.Item.Name = kv[1]
			break
		case "Quantity":
			d.Quantity.Parse(kv[1])
			break
		case "Rarity":
			d.RarityString = kv[1]
			break
		}

	}

	log.Printf("Drop Parsed: %s - %s", d.Item.Name, d.Quantity.String())
	return &d
}
