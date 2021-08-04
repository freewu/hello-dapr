package client

import (
	"github.com/freewu/hello-dapr/common/response"
	"context"
	"encoding/json"
	"fmt"
	"errors"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	log "github.com/sirupsen/logrus"
)

type client struct {
	ServerName string // 服务名称
	daprClient dapr.Client // dapr 请求客户端
}

//@author: [freewu](https://github.com/freewu)
//@function: NewClient
//@description: 通过服务名称获取一个 dapr 请求对象
//@param: serverName string 服务名称
//@return: *client
func NewClient(serverName string) *client {
	c, err := dapr.NewClient()
	if err != nil {
		log.Error(err.Error())
	}
	return &client{ ServerName: serverName,daprClient: c}
}

//@author: [freewu](https://github.com/freewu)
//@function: NewClientWithPort
//@description: 通过服务名称 & 端口号 获取一个 dapr 请求对象
//@param: serverName string 服务名称
//@param: port string 服务端口
//@return: *client
func NewClientWithPort(serverName string,port string) *client {
	c, err := dapr.NewClientWithPort(port)
	if err != nil {
		log.Error(err.Error())
	}
	return &client{ ServerName: serverName,daprClient: c}
}

//@author: [freewu](https://github.com/freewu)
//@function: POST
//@description: post 请求
//@param: url string 方法名称
//@param: data interface{} 请求的数据
//@param: target interface{} 远程调用后得到的数据
//@param: ctx context.Context 上下文信息
//@return: error
func(c *client) POST(url string,data interface{},target interface{},ctx context.Context) error {
	return c.handle(url,data,ctx,http.MethodPost,target)
}

//@author: [freewu](https://github.com/freewu)
//@function: GET
//@description: get 请求
//@param: url string 方法名称
//@param: data interface{} 请求的数据
//@param: target interface{} 远程调用后得到的数据
//@param: ctx context.Context 上下文信息
//@return: error
func(c *client) GET(url string,data interface{},target interface{},ctx context.Context) error {
	return c.handle(url,data,ctx,http.MethodGet,target)
}

//@author: [freewu](https://github.com/freewu)
//@function: PUT
//@description: put 请求
//@param: url string 方法名称
//@param: data interface{} 请求的数据
//@param: target interface{} 远程调用后得到的数据
//@param: ctx context.Context 上下文信息
//@return: error
func(c *client) PUT(url string,data interface{},target interface{},ctx context.Context) error {
	return c.handle(url,data,ctx,http.MethodPut,target)
}

//@author: [freewu](https://github.com/freewu)
//@function: DELETE
//@description: delete 请求
//@param: url string 方法名称
//@param: data interface{} 请求的数据
//@param: target interface{} 远程调用后得到的数据
//@param: ctx context.Context 上下文信息
//@return: error
func(c *client) DELETE(url string,data interface{},target interface{},ctx context.Context) error {
	return c.handle(url,data,ctx,http.MethodDelete,target)
}

//@author: [freewu](http://git.yibianli.com/freewu)
//@function: handle
//@description: 统一的请求处理
//@param: url string 方法名称
//@param: data interface{} 请求的数据
//@param: ctx context.Context 上下文件信息
//@param: method string 请求方法类型
//@param: target interface{} 远程调用后得到的数据
//@return: error
func(c *client) handle(url string,data interface{},ctx context.Context,method string,target interface{}) error {
	value := "application/json"
	// 如果有设置请求类型
	if ctx.Value("ContentType") != nil {
		value = ctx.Value("ContentType").(string)
	}
	bs, _ := json.Marshal(data)
	daprContent := &dapr.DataContent{
		ContentType: value,
		Data: bs,
	}
	resp, err := c.daprClient.InvokeMethodWithContent(ctx, c.ServerName, url, method, daprContent)
	// 请求错误(网络问题)
	if err != nil {
		return err
	}
	// todo 这边全部按自己接口格式封装后期看看这个怎么定制化
	r := &response.Response{}
	err = json.Unmarshal(resp,r)
	// 远程接口的错误
	if r.Code != 0 {
		//return fmt.Errorf("get task fail, code %d, msg:%s", r.Code, r.Msg) // 直接返回远程接口的错误信息
		return errors.New("remote:" + r.Msg)
	}
	// 需要通过 target 把数据传递出去
	if target != nil {
		dataBytes, err := json.Marshal(r.Data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(dataBytes, target)
		if err != nil {
			return fmt.Errorf("get task unmarshal fail, err:%s, origin-data:%s", err.Error(), string(dataBytes))
		}
	}
	return nil
}