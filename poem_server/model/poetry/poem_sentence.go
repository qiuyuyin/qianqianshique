package poetry

import "go.mongodb.org/mongo-driver/bson/primitive"

type PoemSentence struct {
	ID         primitive.ObjectID `bson:"_id"`
	PoemID     primitive.ObjectID `json:"poemID"`
	AuthorID   primitive.ObjectID `json:"authorID"`
	Content    string             `json:"content"`
	Title      string             `json:"title"`
	AuthorName string             `json:"from"`
	PoemType   string             `json:"poemtype"`
	Dynasty    string             `json:"dynasty"`
}
