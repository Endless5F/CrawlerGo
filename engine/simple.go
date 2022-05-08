package engine

import (
	"GoCodeProject/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s", r.Url)

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Panicf("Fetcher: errpr fetching: %s", r.Url)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
