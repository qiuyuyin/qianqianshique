package poem_request

import (
	"poem_server/model/common/request"
	"poem_server/model/poetry"
)

type SearchSentenceParams struct {
	request.PageInfo
	poetry.PoemSentence
}
