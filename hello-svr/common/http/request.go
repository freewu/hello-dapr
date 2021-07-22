package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/service/common"
	"reflect"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// https://github.com/goinggo/mapstructure

// 请求体
type Request struct {
	Context context.Context
	Event *dapr.InvocationEvent
}

func NewRequest(ctx context.Context,event *dapr.InvocationEvent) *Request  {
	return &Request{
		Context: ctx,
		Event: event,
	}
}

var (
	// xmlHeaderBytes is the most common XML format header.
	xmlHeaderBytes = []byte("<?xml")
)

// ParseInvocationData 转换请求参数到结构体
func (d *Request) Parse(pointer interface{}) error {
	var (
		reflectVal1  = reflect.ValueOf(pointer)
		reflectKind1 = reflectVal1.Kind()
	)
	if reflectKind1 != reflect.Ptr {
		return fmt.Errorf(
			"parameter should be type of *struct/**struct/*[]struct/*[]*struct, but got: %v",
			reflectKind1,
		)
	}
	var (
		reflectVal2  = reflectVal1.Elem()
		reflectKind2 = reflectVal2.Kind()
	)

	switch reflectKind2 {
	// Single struct, post content like:
	// 1. {"id":1, "name":"john"}
	// 2. ?id=1&name=john
	case reflect.Ptr, reflect.Struct:
		// 从 url 或者 form 或者 body 获取 data
		data := d.Form()

		err := gconv.Struct(data, pointer)
		if err != nil {
			return err
		}

		//// Validation.
		//if err := gvalid.CheckStruct(pointer, nil); err != nil {
		//	return err
		//}

	// Multiple struct, it only supports JSON type post content like:
	// [{"id":1, "name":"john"}, {"id":, "name":"smith"}]
	case reflect.Array, reflect.Slice:
		// If struct slice conversion, it might post JSON/XML content,
		// so it uses gjson for the conversion.
		j, err := gjson.LoadContent(d.Event.Data)
		if err != nil {
			return err
		}
		if err := j.GetStructs(".", pointer); err != nil {
			return err
		}
		//for i := 0; i < reflectVal2.Len(); i++ {
		//	if err := gvalid.CheckStruct(reflectVal2.Index(i), nil); err != nil {
		//		return err
		//	}
		//}
	}
	return nil
}

func (d *Request) Form() map[string]interface{} {
	m := make(map[string]interface{})

	// url query
	queryMap, _ := gstr.Parse(d.Event.QueryString)

	// form
	formMap := make(map[string]interface{})
	if gstr.Contains(d.Event.ContentType, "form") {
		formMap, _ = gstr.Parse(string(d.Event.Data))
	}

	// body
	bodyMap := make(map[string]interface{})
	if body := d.Event.Data; len(body) > 0 {
		body = bytes.TrimSpace(body)

		// json
		if body[0] == '{' && body[len(body)-1] == '}' {
			json.Unmarshal(body, &bodyMap)
		}

		// XML
		if len(body) > 5 && bytes.EqualFold(body[:5], xmlHeaderBytes) {
			bodyMap, _ = gxml.DecodeWithoutRoot(body)
		}
		if body[0] == '<' && body[len(body)-1] == '>' {
			bodyMap, _ = gxml.DecodeWithoutRoot(body)
		}

		// Default parameters decoding.
		if bodyMap == nil {
			bodyMap, _ = gstr.Parse(gconv.UnsafeBytesToStr(d.Event.Data))
		}
	}

	// 统一到 map
	for key, value := range queryMap {
		m[key] = value
	}

	for key, value := range formMap {
		m[key] = value
	}

	for key, value := range bodyMap {
		m[key] = value
	}

	return m
}

func (d *Request) Ints(key string) []int {
	dataMap := d.Form()
	return gvar.New(dataMap[key]).Ints()
}

func (d *Request) UInts(key string) []uint {
	dataMap := d.Form()
	return gvar.New(dataMap[key]).Uints()
}

func (d *Request) GetDataUInts32(key string) []uint32 {
	values := d.UInts(key)
	result := []uint32{}

	for _, v := range values {
		result = append(result, uint32(v))
	}
	return result
}