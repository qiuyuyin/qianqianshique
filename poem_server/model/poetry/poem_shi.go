package poetry

import "go.mongodb.org/mongo-driver/bson/primitive"

type PoemType struct {
	ID           primitive.ObjectID `bson:"_id"`
	Title        string             `json:"title"`
	AuthorName   string             `json:"authorName"`
	AuthorID     string             `json:"authorId"`
	Dynasty      string             `json:"dynasty"`
	Appreciation []string           `json:"appreciation"`
	Content      []string           `json:"content"`
	Comment      []string           `json:"comment"`
	Translation  []string           `json:"translation"`
	Intro        string             `json:"intro"`
	Annotation   []string           `json:"annotation"`
	Rank         uint               `json:"rank"`
	PoemType     string             `json:"poemType" binding:"required"`
	Tags         []string           `json:"tags"`
	Tag          string             `json:"tag"`
	PoemStyle    string             `json:"poemStyle"`
}
