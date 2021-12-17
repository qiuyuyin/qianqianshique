package poetry

// 注意两个request包的不同使用方式
import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"poem_server/global"
	"poem_server/model/common/responce"
	"poem_server/model/poetry/poem_request"
	"poem_server/utils"
)

type PoemApi struct {
}

// @Tags PoemApi（POST进行请求要将ID删去）
// @Summary 分页获取Poem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body poem_request.SearchPoemParams true "分页获取Author列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /poem/getPoemList [post]
func (a *PoemApi) GetPoemList(c *gin.Context) {
	var pageInfo poem_request.SearchPoemParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil || !utils.PoemTypeRight(pageInfo.PoemType.PoemType) {
		responce.FailWithMessage("参数不正确", c)
	} else if err, list, count := poemService.GetPoemByTagAndPage(pageInfo); err != nil {
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

// @Tags PoemApi（POST进行请求要将ID删去）
// @Summary 根据OID获取Poem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param oid path string true "oid"
// @Param poemType path string true "poemType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /poem/getPoemByID/{poemType}/{oid} [get]
func (a *PoemApi) GetPoemById(c *gin.Context) {
	oid := c.Param("oid")
	poemType := c.Param("poemType")
	if oid == "" || !utils.PoemTypeRight(poemType) {
		responce.FailWithMessage("提供的参数格式不正确", c)
	} else if err, poem := poemService.GetPoemByID(oid, poemType); err != nil {
		global.POEM_LOG.Error("获取失败", zap.Any("err", err))
		responce.FailWithMessage("获取失败", c)
	} else {
		responce.OkWithDetailed(poem, "获取作者成功", c)
	}
}
