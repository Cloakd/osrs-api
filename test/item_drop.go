package main

import (
	"log"

	osrs_api "github.com/cloakd/osrs-api"
)

func main() {
	str := "1-4&lt;!-- This is extra shard drops; you can get 5 total, including the guaranteed 1. The extra shards are max 4. -->"

	q := osrs_api.Quantity{}
	q.Parse(str)

	log.Printf(q.String())
}
