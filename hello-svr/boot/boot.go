package boot

import (
	"demo/hello-svr/handler"
	log "github.com/sirupsen/logrus"
	daprd "github.com/dapr/go-sdk/service/grpc"
)

func Init() error {
	log.Println("hello-svr starting...")

	// 注册 dapr 服务
	service, err := daprd.NewService(":9001")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
		return err
	}

	// 注册 http 路由
	if err := handler.InitHttp(service); err != nil {
		log.Fatalf("handler register failed: %v", err)
		return err
	}

	// 启动服务
	if err := service.Start(); err != nil {
		log.Fatalf("error listenning: %v", err)
		return err
	}
	return nil
}