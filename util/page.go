package util

import (
	"math"
	"strconv"
)

type Page struct {
	PageSize  int64       `json:"pageSize"`
	PageNo    int64       `json:"pageNo"`
	Total     int64       `json:"total"`
	TotalPage int64       `json:"totalPage"`
	Result    interface{} `json:"result"`
	Version   string      `json:"version"`
}

func NewPage(pageNo, pageSize int64) Page {
	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	if pageSize > 150 {
		pageSize = 150
	}

	return Page{
		PageSize: pageSize,
		PageNo:   pageNo,
	}
}

func NewPageWithStr(pageNoStr, pageSizeStr string) Page {
	pageNo, err := strconv.ParseInt(pageNoStr, 10, 64)
	if err != nil || pageNo < 1 {
		pageNo = 1
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	if pageSize > 150 {
		pageSize = 150
	}

	return Page{
		PageSize: pageSize,
		PageNo:   pageNo,
	}
}

func (p Page) SetTotal(total int64) Page {
	if p.PageSize == 0 {
		p.PageSize = 10
	}

	p.Total = total
	p.TotalPage = int64(math.Ceil(float64(p.Total) / float64(p.PageSize)))

	return p
}

func (p Page) Offset() int {
	offset := (p.PageNo - 1) * p.PageSize
	if offset < 0 {
		return 0
	}
	return int(offset)
}
