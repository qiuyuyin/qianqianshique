package request

// PageInfo 分页查询的信息
type PageInfo struct {
	Page     int `json:"page" form:"page" binding:"required,gt=0"`               // 页码
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,gt=4,lt=12"` // 每页大小
}
