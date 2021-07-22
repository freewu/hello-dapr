package handler

import (
	"demo/hello-svr/common/http"
)

// 用来处理验证 restful 接口
type Restful struct {
}

// @Summary Restful-Get
// @Description Restful-Get
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /restful/get [get]
// @Security
func (c *Restful) Get(request *http.Request) string {
	return "get"
}

// @Summary Restful-Post
// @Description Restful-Post
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /restful/post [post]
// @Security
func (c *Restful) Post(request *http.Request) string {
	return "post"
}

// @Summary Restful-Put
// @Description Restful-Put
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /restful/put [put]
// @Security
func (c *Restful) Put(request *http.Request) string {
	return "put"
}

// @Summary Restful-Delete
// @Description Restful-Delete
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /restful/delete [delete]
// @Security
func (c *Restful) Delete(request *http.Request) string {
	return "delete"
}