package poem_responce

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchPoemType struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title"`
	AuthorName string             `json:"authorName"`
	Dynasty    string             `json:"dynasty"`
	PoemType   string             `json:"poemType" binding:"required"`
}

type SearchAuthorType struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `json:"name" `
	Dynasty string             `json:"dynasty"`
}

type SearchSentenceType struct {
	ID       primitive.ObjectID `bson:"_id"`
	PoemID   primitive.ObjectID `json:"poemID"`
	AuthorID primitive.ObjectID `json:"authorID"`
	Content  string             `json:"content"`
	PoemType string             `json:"poemType"`
}

type SearchResponse struct {
	SearchPoemList     []SearchPoemType
	SearchAuthorList   []SearchAuthorType
	SearchSentenceList []SearchSentenceType
}
