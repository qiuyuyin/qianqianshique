package poetry

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"poem_server/global"
	"poem_server/model/common/responce"
	"poem_server/model/poetry/poem_request"
	"strconv"
)

type AuthorApi struct {
}

// @Tags AuthorApi（POST进行请求要将ID删去）
// @Summary 分页获取Author列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body poem_request.SearchAuthorParams true "分页获取Author列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /author/getAuthorList [post]
func (a *AuthorApi) GetAuthorList(c *gin.Context) {
	var pageInfo poem_request.SearchAuthorParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		responce.FailWithMessage("参数不正确", c)
	} else if err, list, count := authorService.GetAuthorInfoList(pageInfo); err != nil {
		global.POEM_LOG.Error("获取失败", zap.Any("err", err))
		responce.FailWithMessage("获取失败", c)
	} else {
		responce.OkWithDetailed(responce.PageResult{
			List:     list,
			Total:    count,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags AuthorApi（POST进行请求要将ID删去）
// @Summary 根据OID获取Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param oid path string true "oid"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /author/getAuthorByID/{oid} [get]
func (a *AuthorApi) GetAuthorById(c *gin.Context) {
	oid := c.Param("oid")
	if err, author := authorService.GetAuthorByID(oid); err != nil {
		global.POEM_LOG.Error("获取失败", zap.Any("err", err))
		responce.FailWithMessage("获取失败", c)
	} else {
		responce.OkWithDetailed(author, "获取作者成功", c)
	}
}

func (a *AuthorApi) GetAuthorPoemCountByLimit(c *gin.Context) {
	limit := c.Param("limit")
	limitParam, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		responce.FailWithMessage("参数错误", c)
		return
	}
	if err, author := authorService.GetAuthorPoemCountByLimit(uint(limitParam)); err != nil {
		global.POEM_LOG.Error("获取失败", zap.Any("err", err))
		responce.FailWithMessage("获取失败", c)
	} else {
		responce.OkWithDetailed(author, "获取作者成功", c)
	}
}
