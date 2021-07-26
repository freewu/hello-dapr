package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page"`         // 页码
	PageSize int `json:"pageSize"` // 每页大小
}

type SortInfo struct {
	OrderBy string `json:"orderBy"`// 排序
}

// Find by id structure
type GetById struct {
	ID uint `json:"id"` // 主键ID
}

type IdsReq struct {
	Ids []string `json:"ids"`
}

type Empty struct{}