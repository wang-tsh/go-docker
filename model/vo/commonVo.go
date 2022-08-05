package vo

import "container/list"

type BaseResultVo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Page struct {
	Records list.List
	Total   int32
	Size    int
	Current int
}
