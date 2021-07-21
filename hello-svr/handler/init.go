package handler

import (
	dapr "github.com/dapr/go-sdk/service/common"
	log "github.com/sirupsen/logrus"
)

// http 路由注册
func InitHttp(s dapr.Service) error {
	// todo 想法是 xxxx.GET("/hello",(&Greeting{}).Hello)
	log.Println("http router register")
	// 注册路由
	if err := s.AddServiceInvocationHandler("hello", (&Greeting{}).Hello); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
		return err
	}

	return nil
}