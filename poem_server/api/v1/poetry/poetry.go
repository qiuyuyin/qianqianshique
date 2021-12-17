package poetry

import "poem_server/service"

type PoemApiGroup struct {
	AuthorApi
	PoemApi
	SentenceApi
	SearchApi
}

var authorService = service.PoemServiceGroupApp.PoetryServiceGroup.AuthorService
var poemService = service.PoemServiceGroupApp.PoetryServiceGroup.PoemService
var sentenceService = service.PoemServiceGroupApp.PoetryServiceGroup.SentenceService
var searchService = service.PoemServiceGroupApp.PoetryServiceGroup.SearchService
