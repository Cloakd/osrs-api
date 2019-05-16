package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	osrs_api "github.com/cloakd/osrs-api"
	"github.com/cloakd/osrs-api/cmd"
)

const (
	ITEM_CACHE_FILE = "cache/item_cache.json"
)

func main() {
	start := time.Now()

	monMap, err := cmd.CacheItems()
	if err != nil {
		log.Fatal(err)
	}

	mWrite := make(chan *osrs_api.Item, 100)
	go func(mChan chan *osrs_api.Item) {
		log.Printf("Opening file: %s", ITEM_CACHE_FILE)
		f, err := os.Create(ITEM_CACHE_FILE)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		for m := range mChan {
			j, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}

			_, err = f.Write(j)
			if err != nil {
				log.Fatal(err)
			}

			f.WriteString("\n")
		}
	}(mWrite)

	ms := cmd.NewItemScraper(&monMap)
	ms.Scrape(mWrite)

	close(mWrite)

	fin := time.Since(start)
	log.Printf("Item scrape took %s", fin)
}
