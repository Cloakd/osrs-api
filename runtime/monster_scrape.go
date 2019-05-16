package main

import "github.com/cloakd/osrs-api/cmd"

func main() {
	ms := cmd.MonsterScraper{}

	ms.Scrape()
}
