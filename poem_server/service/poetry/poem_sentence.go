package poetry

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"poem_server/global"
	"poem_server/model/poetry"
	"poem_server/model/poetry/poem_request"
	"poem_server/utils"
)

type SentenceService struct {
}

func (s SentenceService) GetSentenceByPage(info poem_request.SearchSentenceParams) (err error, list interface{}, count int64) {
	limit := info.PageSize
	offset := limit * (info.Page - 1)
	coll := global.POEM_MONGO.Collection("sentence")
	ctx := context.Background()
	var sentenceList []poetry.PoemSentence
	filter := bson.M{}
	m := map[string]interface{}{}
	if info.PoemSentence.Content != "" {
		m["content"] = bson.M{"$regex": info.PoemSentence.Content}
	}
	if info.PoemSentence.Dynasty != "" {
		m["dynasty"] = bson.M{"$regex": info.PoemSentence.Dynasty}
	}
	if info.PoemSentence.PoemType != "" {
		m["poemtype"] = info.PoemSentence.PoemType
	}
	if err, filter = utils.ParseMapToBson(m); err != nil {
		return err, nil, 0
	}
	count, err = coll.Find(ctx, filter).Count()
	if err != nil {
		return err, sentenceList, count
	} else {
		err := coll.Find(ctx, filter).Skip(int64(offset)).Limit(int64(limit)).All(&sentenceList)
		if err != nil {
			return err, sentenceList, count
		}
	}
	return err, sentenceList, count
}
