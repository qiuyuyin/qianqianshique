package poetry

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"poem_server/global"
	"poem_server/model/poetry/poem_responce"
)

type SearchService struct {
}

func (s SearchService) SearchAll(keyword string) (error, poem_responce.SearchResponse) {
	var res poem_responce.SearchResponse
	collection := global.POEM_MONGO.Collection("author")
	ctx := context.Background()
	var authorList []poem_responce.SearchAuthorType
	err := collection.Find(ctx, bson.M{"name": bson.M{"$regex": keyword}}).Limit(int64(3)).All(&authorList)
	if err != nil {
		return err, poem_responce.SearchResponse{}
	}
	res.SearchAuthorList = authorList

	collection = global.POEM_MONGO.Collection("poem_shi")
	var poemList []poem_responce.SearchPoemType

	err = collection.Find(ctx, bson.M{"title": bson.M{"$regex": keyword}}).Limit(int64(4)).All(&poemList)
	if err != nil {
		return err, poem_responce.SearchResponse{}
	}
	for i := range poemList {
		poemList[i].PoemType = "shi"
	}
	res.SearchPoemList = poemList

	collection = global.POEM_MONGO.Collection("sentence")
	var sentenceList []poem_responce.SearchSentenceType
	err = collection.Find(ctx, bson.M{"content": bson.M{"$regex": keyword}}).Limit(int64(2)).All(&sentenceList)
	if err != nil {
		return err, poem_responce.SearchResponse{}
	}
	res.SearchSentenceList = sentenceList
	if len(sentenceList) == 0 && len(poemList) == 0 && len(authorList) == 0 {
		return errors.New("没有查询出来什么"), res
	}
	return err, res
}
