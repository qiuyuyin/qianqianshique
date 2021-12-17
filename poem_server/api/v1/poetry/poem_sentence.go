package poetry

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"poem_server/global"
	"poem_server/model/common/responce"
	"poem_server/model/poetry/poem_request"
)

type SentenceApi struct {
}

// @Tags SentenceApi（POST进行请求要将ID删去）
// @Summary 分页获取Sentence列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body poem_request.SearchSentenceParams true "分页获取Sentence列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sentence/getSentenceList [post]
func (a *SentenceApi) GetSentenceList(c *gin.Context) {
	var sentencePageInfo poem_request.SearchSentenceParams
	err := c.ShouldBindJSON(&sentencePageInfo)
	if err != nil {
		responce.FailWithMessage("参数不正确", c)
	} else if err, list, count := sentenceService.GetSentenceByPage(sentencePageInfo); err != nil {
		global.POEM_LOG.Error("获取失败", zap.Any("err", err))
		responce.FailWithMessage("获取失败", c)
	} else {
		responce.OkWithDetailed(responce.PageResult{
			List:     list,
			Total:    count,
			Page:     sentencePageInfo.Page,
			PageSize: sentencePageInfo.PageSize,
		}, "获取成功", c)
	}
}
