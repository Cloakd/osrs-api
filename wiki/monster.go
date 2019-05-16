package wiki

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type (
	MonsterIndex struct {
		IndexPages []string

		monsterLinks sync.Map
		wait         sync.WaitGroup
	}
)

func GetMonsters() []string {
	mi := MonsterIndex{
		monsterLinks: sync.Map{},
		IndexPages: []string{
			"Bestiary/Levels_1_to_10",
			"Bestiary/Levels_11_to_20",
			"Bestiary/Levels_21_to_30",
			"Bestiary/Levels_31_to_40",
			"Bestiary/Levels_41_to_50",
			"Bestiary/Levels_51_to_60",
			"Bestiary/Levels_61_to_70",
			"Bestiary/Levels_71_to_80",
			"Bestiary/Levels_81_to_90",
			"Bestiary/Levels_91_to_100",
			"Bestiary/Levels_higher_than_100",
		},
	}

	mi.ExpandIndex()

	return mi.BuildIndex()
}

//Expand out the index pages into all monster links¬
func (i *MonsterIndex) ExpandIndex() {
	log.Printf("Expanding Page Indexes")

	i.wait = sync.WaitGroup{}
	for _, index := range i.IndexPages {
		go i.links(index, &i.monsterLinks)
	}
	i.wait.Wait()
}

func (i *MonsterIndex) BuildIndex() []string {
	log.Printf("Building Page Indexes")

	i.wait = sync.WaitGroup{}

	var idx []string
	i.monsterLinks.Range(func(key, value interface{}) bool {
		log.Printf("%v - %s", key, value)

		idx = append(idx, value.(string))
		return true
	})

	return idx
}

func (i *MonsterIndex) links(index string, pageMap *sync.Map) {
	i.wait.Add(1)
	defer i.wait.Done()

	url := fmt.Sprintf(WIKI_EDIT, index)
	log.Printf("Scraping index links: %s (%s)", index, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		//return nil, err
	}

	defer resp.Body.Close()
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		//return nil, err
	}

	//log.Printf("%s", str)

	//return str, nil
	p := NewParser(str)
	links := p.Links()

	for _, l := range links {
		if parsed, ok := i.parseLink(l); ok {
			log.Printf("Storing link: %s (%v_%v)", parsed, index, parsed)

			pageMap.Store(fmt.Sprintf("%v_%v", index, parsed), parsed)
		}
	}
}

func (i *MonsterIndex) parseLink(link string) (string, bool) {
	if strings.Contains(link, ":") {
		return "", false
	}

	link = strings.ReplaceAll(link, "[[", "")
	link = strings.ReplaceAll(link, "]]", "")
	link = strings.ReplaceAll(link, " ", "_")

	return link, true
}
