package poetry

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"poem_server/global"
	"poem_server/model/poetry"
	"poem_server/model/poetry/poem_request"
	"poem_server/model/poetry/poem_responce"
	"poem_server/utils"
)

type PoemService struct {
}

func (poemService PoemService) GetPoemByTagAndPage(info poem_request.SearchPoemParams) (err error, list interface{}, count int64) {
	limit := info.PageSize
	offset := limit * (info.Page - 1)
	log.Println(limit)
	log.Println(offset)
	// 根据用户输入的类型来进行查找对应类型的数据。因为和数据集相关联所以这个类型数据必须存在
	coll := global.POEM_MONGO.Collection("poem_" + info.PoemType.PoemType)
	ctx := context.Background()
	// 声明了poemList类型
	var poemList []poem_responce.PoemType
	// 定义过滤器类型
	filter := bson.M{}
	m := map[string]interface{}{}
	if info.PoemType.AuthorName != "" {
		m["authorName"] = bson.M{"$regex": info.PoemType.AuthorName}
	}
	if info.PoemType.Title != "" {
		m["title"] = bson.M{"$regex": info.PoemType.Title}
	}
	if info.PoemType.Dynasty != "" {
		m["dynasty"] = info.PoemType.Dynasty
	}
	if info.PoemType.PoemStyle != "" {
		m["poemStyle"] = info.PoemType.PoemStyle
	}
	if info.PoemType.Tag != "" {
		m["tags"] = info.PoemType.Tag
	}
	if err, filter = utils.ParseMapToBson(m); err != nil {
		return err, nil, 0
	}
	// 找到总体的数量，为分页做准备
	count, err = coll.Find(ctx, filter).Count()
	if err != nil {
		return err, poemList, count
	} else {
		// 通过rank来进行查找
		err := coll.Find(ctx, filter).Sort("-rank").Skip(int64(offset)).Limit(int64(limit)).All(&poemList)
		if err != nil {
			return err, poemList, count
		}
	}
	for i := range poemList {
		poemList[i].PoemType = info.PoemType.PoemType
	}
	return err, poemList, count
}

// 根据ID和诗句类型获取诗句
func (poemService PoemService) GetPoemByID(ID string, PoemType string) (err error, list interface{}) {
	coll := global.POEM_MONGO.Collection("poem_" + PoemType)
	ctx := context.Background()
	var poem poetry.PoemType
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err, poem
	} else {
		err := coll.Find(ctx, bson.M{"_id": oid}).One(&poem)
		if err != nil {
			return err, poem
		}
	}
	return err, poem
}
