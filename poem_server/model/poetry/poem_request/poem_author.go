package poem_request

import (
	"poem_server/model/common/request"
	"poem_server/model/poetry"
)

// author 查询结构体
type SearchAuthorParams struct {
	request.PageInfo
	poetry.PoemAuthor
}
