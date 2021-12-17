package poetry

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"poem_server/global"
	"poem_server/model/poetry"
	"poem_server/model/poetry/poem_request"
	"poem_server/model/poetry/poem_responce"
	"poem_server/utils"
)

type AuthorService struct {
}

func (authorService AuthorService) GetAuthorInfoList(info poem_request.SearchAuthorParams) (err error, list interface{}, count int64) {
	limit := info.PageSize
	offset := limit * (info.Page - 1)
	coll := global.POEM_MONGO.Collection("author")
	ctx := context.Background()
	var authorList []poem_responce.PoemAuthor
	// 定义过滤器类型
	filter := bson.M{}
	m := map[string]interface{}{}
	if info.PoemAuthor.Name != "" {
		m["name"] = bson.M{"$regex": info.PoemAuthor.Name}
	}
	if info.PoemAuthor.Dynasty != "" {
		m["dynasty"] = bson.M{"$regex": info.PoemAuthor.Dynasty}
	}
	if err, filter = utils.ParseMapToBson(m); err != nil {
		return err, nil, 0
	}

	count, err = coll.Find(ctx, bson.M{}).Count()
	if err != nil {
		return err, authorList, count
	} else {
		err := coll.Find(ctx, filter).Sort("-quantity").Skip(int64(offset)).Limit(int64(limit)).All(&authorList)
		if err != nil {
			return err, authorList, count
		}
	}
	return err, authorList, count
}

func (authorService AuthorService) GetAuthorByID(ID string) (err error, list interface{}) {
	coll := global.POEM_MONGO.Collection("author")
	ctx := context.Background()
	var author poetry.PoemAuthor
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err, author
	} else {
		err := coll.Find(ctx, bson.M{"_id": oid}).One(&author)
		if err != nil {
			return err, author
		}
	}
	return err, author
}

func (authorService AuthorService) GetAuthorPoemCountByLimit(limit uint) (err error, list interface{}) {
	coll := global.POEM_MONGO.Collection("author")
	ctx := context.Background()
	var authorList []poem_responce.PoemCount
	// 查找类型
	err = coll.Find(ctx, bson.M{}).Sort("-poemCount").Limit(int64(limit)).All(&authorList)
	if err != nil {
		return err, nil
	}
	return err, authorList
}
