package osrs_api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DropCollection struct {
	Drops []*ItemDrop
}

func (dc *DropCollection) Add(drop *ItemDrop) {
	dc.Drops = append(dc.Drops, drop)
}

type ItemDrop struct {
	Item         Item
	ItemId       int32
	Quantity     Quantity
	Rarity       float32
	RarityString string
	RarityNotes  []*RarityNote
}

type Quantity struct {
	Min   int
	Max   int
	Noted bool
}

func (q *Quantity) Parse(str string) {

	if strings.Contains(str, "noted") || strings.Contains(str, "Noted") {
		str = strings.ReplaceAll(str, "(noted)", "")
		str = strings.ReplaceAll(str, "(Noted)", "")
		q.Noted = true
	}

	if "" == strings.ReplaceAll(str, " ", "") {
		q.Min = 1
		q.Max = 1
		return
	}

	//Fix up common issues
	str = strings.ReplaceAll(str, ",", "-")
	str = strings.ReplaceAll(str, ";", "-")

	//Remove descriptions on items
	if strings.Contains(str, "&lt") {
		split := strings.Split(str, "&lt")
		str = split[0] //Only want the first part
	}

	res, err := strconv.Atoi(strings.ReplaceAll(str, " ", ""))
	if err == nil {
		q.Min = res
		q.Max = res
		return
	}

	slice := strings.Split(str, "-")
	if len(slice) >= 2 {
		var err error

		min := strings.ReplaceAll(slice[0], " ", "")
		max := strings.ReplaceAll(slice[len(slice)-1], " ", "")

		q.Min, err = strconv.Atoi(min)
		if err != nil {
			log.Fatal(err)
		}

		q.Max, err = strconv.Atoi(max)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	log.Fatalf("Unable to parse quantity string: %s", str)
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
