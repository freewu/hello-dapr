package handler

import (
	"github.com/freewu/hello-dapr/common/http"
	dapr "github.com/dapr/go-sdk/service/common"
	log "github.com/sirupsen/logrus"
)

// http 路由注册
func InitHttp(s dapr.Service) (err error) {
	// todo 想法是 xxxx.GET("/hello",(&Greeting{}).Hello)
	log.Println("http router register")
	register := http.NewRouteRegister(s)

	// 注册路由
	err = register.GET("hello", (&Greeting{}).Hello)
	//if err := s.AddServiceInvocationHandler("hello", (&Greeting{}).Hello); err != nil {
	//	log.Fatalf("error adding invocation handler: %v", err)
	//	return err
	//}

	// restful
	rest := &Restful{}
	err = register.GET("restful/get", rest.Get)
	err = register.POST("restful/post", rest.Post)
	err = register.PUT("restful/put", rest.Put)
	err = register.DELETE("restful/delete", rest.Delete)

	// params
	params := &Params{}
	err = register.GET("params/get", params.Get)
	err = register.POST("params/body", params.Body)
	err = register.POST("params/form", params.Form)

	return err
}