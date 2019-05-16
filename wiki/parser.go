package wiki

import (
	"fmt"
	"log"
	"regexp"
)

type Parser struct {
	Body []byte
}

const (
	LINK_REGEX = `\[\[(.*?)\]\]`
	TAG_REGEX  = `\{\{%s(.*?)\}\}`
)

func NewParser(resp []byte) *Parser {
	return &Parser{
		Body: resp,
	}
}

//Returns array of strings
func (p *Parser) Tags(tag string) [][]byte {
	pattern := fmt.Sprintf(TAG_REGEX, tag)
	//log.Printf("Using pattern %s", pattern)

	re := regexp.MustCompile(pattern)
	return re.FindAll(p.Body, -1)
}

func (p *Parser) Links() []string {
	//log.Printf("Using pattern %s", LINK_REGEX)

	re := regexp.MustCompile(LINK_REGEX)

	results := re.FindAll(p.Body, -1)
	log.Printf("Found %v links", len(results))

	output := make([]string, len(results))
	for i, v := range results {
		output[i] = string(v)
	}

	return output
}
