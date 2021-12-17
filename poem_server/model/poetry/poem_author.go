package poetry

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PoemAuthor struct {
	ID        primitive.ObjectID `bson:"_id"`
	Avatar    string             `json:"avatar"`
	Describe  []Describe         `json:"describe"`
	OnlyID    string             `json:"onlyId"`
	Name      string             `json:"name" `
	Dynasty   string             `json:"dynasty"`
	Quantity  int                `json:"quantity"`
	Lifetime  string             `json:"lifetime"`
	UpdateAt  time.Time          `json:"updateAt"`
	AuthorId  uint               `bson:"authorId"`
	ShiCount  uint               `json:"shiCount"`
	QuCount   uint               `json:"quCount"`
	CiCount   uint               `json:"ciCount"`
	PoemCount uint               `json:"poemCount"`
}

type Describe struct {
	Type    string        `json:"type" gson:"type"`
	Content []string      `json:"content" gson:"content"`
	Res     []interface{} `json:"res" `
}
