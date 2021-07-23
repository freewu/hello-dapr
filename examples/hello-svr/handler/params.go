package handler

import (
	"github.com/freewu/hello-dapr/common/http"
	"github.com/freewu/hello-dapr/common/response"
	"demo/hello-svr/model"
)

type Params struct {
}

// @Summary 获取 QueryString 参数
// @Description 从 QueryString 获取参数
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /params/get [get]
// @Security
func (c *Params) Get(request *http.Request) interface{} {
	var user *model.User
	if err := request.Parse(&user); err != nil {
		return response.Fail(-1, err.Error())
	}
	return response.OK(user)
}

// @Summary 获取 Body/Form 参数
// @Description 从 Body/Form 获取参数
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /params/form [POST]
// @Security
func (c *Params) Form(request *http.Request) interface{} {
	var user *model.User
	if err := request.Parse(&user); err != nil {
		return response.Fail(-1, err.Error())
	}
	return response.OK(user)
}

// @Summary 获取 Body/json 参数
// @Description 从 Body/json 获取参数
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /params/body [POST]
// @Security
func (c *Params) Body(request *http.Request) interface{} {
	var user *model.User
	if err := request.Parse(&user); err != nil {
		return response.Fail(-1, err.Error())
	}
	return response.OK(user)
}