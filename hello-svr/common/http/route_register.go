package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	dapr "github.com/dapr/go-sdk/service/common"
	log "github.com/sirupsen/logrus"
	"reflect"
)

// 注册路由相关
type RouteRegister struct {
	Service dapr.Service
}

func NewRouteRegister(service dapr.Service) *RouteRegister {
	return &RouteRegister{
		Service: service,
	}
}

func(r *RouteRegister) GET(path string, object interface{}) {
	// 匿名函数处理
	err := r.Service.AddServiceInvocationHandler(path, func(ctx context.Context, in *dapr.InvocationEvent) (out *dapr.Content, err error) {
		return r.handle("GET",object,ctx,in)
	})
	if err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
}

func(r *RouteRegister) PUT(path string, object interface{}) {
	// 匿名函数处理
	err := r.Service.AddServiceInvocationHandler(path, func(ctx context.Context, in *dapr.InvocationEvent) (out *dapr.Content, err error) {
		return r.handle("PUT",object,ctx,in)
	})
	if err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
}

func(r *RouteRegister) POST(path string, object interface{}) {
	// 匿名函数处理
	err := r.Service.AddServiceInvocationHandler(path, func(ctx context.Context, in *dapr.InvocationEvent) (out *dapr.Content, err error) {
		return r.handle("POST",object,ctx,in)
	})
	if err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
}

func(r *RouteRegister) DELETE(path string, object interface{}) {
	// 匿名函数处理
	err := r.Service.AddServiceInvocationHandler(path, func(ctx context.Context, in *dapr.InvocationEvent) (out *dapr.Content, err error) {
		return r.handle("DELETE",object,ctx,in)
	})
	if err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
}

func(r *RouteRegister) handle(method string,object interface{},ctx context.Context, in *dapr.InvocationEvent) (out *dapr.Content, err error) {
	if in == nil {
		return nil, errors.New("invocation parameter required")
	}
	if in.Verb != method {
		return nil, errors.New("no such request type, please use " + method + " Request")
	}
	// 真正的业务逻辑实现
	function := reflect.ValueOf(object)
	//判断是否是方法对象
	if function.Kind() != reflect.Func {
		return nil, errors.New("it is no a func to invoke")
	}
	// 执行传入方法
	params := make([]reflect.Value,1)  //参数
	params[0] = reflect.ValueOf(in)
	re := function.Call(params)

	var bytes []byte
	// 判断返回类型是否是 string 模板
	if "string" == string(re[0].Type().String()) {
		s := fmt.Sprintf("%v", re[0].Interface())
		bytes = []byte(s)
	} else { // 如果不是string类型直接使用 json 序列化返回
		bytes, err = json.Marshal(re[0].Interface())
		if err != nil {
			return nil,err
		}
	}

	return &dapr.Content{
		Data: bytes,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}, nil
}