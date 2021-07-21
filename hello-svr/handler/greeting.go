package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dapr/go-sdk/service/common"
	log "github.com/sirupsen/logrus"
)

type Greeting struct {
}

// @Summary hello
// @Description hello
// @Tags 测试接口
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /greeting/hello [get]
// @Security
func (c *Greeting) Hello(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	log.Print()
	// r *ghttp.Request
	//ctx
	if in == nil {
		return nil, errors.New("invocation parameter required")
	}
	if in.Verb != "GET" {
		return nil, errors.New("no such request type, please use Get Request")
	}

	aaa := make(map[string]interface{})
	aaa["string"] = "Hello"
	data, _ := json.Marshal(aaa)

	return &common.Content{
		Data:        data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}, nil
}