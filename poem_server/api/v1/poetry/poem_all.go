package poetry

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"poem_server/global"
	"poem_server/model/common/responce"
	"regexp"
)

type SearchApi struct {
}

func (s SearchApi) SearchAll(c *gin.Context) {
	keyword, _ := c.GetQuery("keyword")
	// 必须要存在一个或者多个中文字符，不可存在数字和英文
	compile, _ := regexp.Compile("^[\u4e00-\u9fa5]+$")
	if !compile.MatchString(keyword) {
		responce.FailWithMessage("获取失败", c)
		return
	}
	if err, res := searchService.SearchAll(keyword); err != nil {
		global.POEM_LOG.Error("获取失败", zap.Any("err", err))
		responce.FailWithMessage("获取失败", c)
	} else {
		responce.OkWithDetailed(res, "查询成功", c)
	}
}
