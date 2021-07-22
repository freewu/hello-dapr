package http

import (
	"context"
	dapr "github.com/dapr/go-sdk/service/common"

)

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
