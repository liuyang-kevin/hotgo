package form

import (
	"hotgo/internal/consts"
)

type ReqPageFunc interface {
	GetPage() int
	GetPerPage() int
}

// PageReq 分页请求
type PageReq struct {
	Page    int `json:"page" example:"10" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PerPage int `json:"pageSize" example:"1" d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}

// GetPage 获取当前页码
func (req *PageReq) GetPage() int {
	return req.Page
}

// GetPerPage 获取每页数量
func (req *PageReq) GetPerPage() int {
	return req.PerPage
}

// PageRes 分页响应
type PageRes struct {
	PageReq
	PageCount int `json:"pageCount" example:"0" dc:"全部数据量"`
}

// Pack 打包分页数据
func (res *PageRes) Pack(req ReqPageFunc, totalCount int) {
	res.PageCount = CalPageCount(totalCount, req.GetPerPage())
	res.Page = req.GetPage()
	res.PerPage = req.GetPerPage()
}

func CalPageCount(totalCount int, perPage int) int {
	return (totalCount + perPage - 1) / perPage
}

// CalPage 计算分页偏移量
func CalPage(page, perPage int) (newPage, newPerPage int, offset int) {
	if page <= 0 {
		newPage = consts.DefaultPage
	} else {
		newPage = page
	}
	if perPage <= 0 {
		newPerPage = consts.DefaultPageSize
	} else {
		newPerPage = perPage
	}

	offset = (newPage - 1) * newPerPage
	return
}
