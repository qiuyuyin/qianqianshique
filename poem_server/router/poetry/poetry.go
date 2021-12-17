package poetry

type RouterGroup struct {
	AuthorRouter
	PoemRouter
	SearchRouter
	SentenceRouter
}
