SRV=demo-hello-srv
VERSION=latest
APP_PORT=9001 # 服务所占用的端口
DAPR_HTTP_PORT=9101 # 对外提供的 http 服务端口
DAPR_GRPC_PORT=9201 # 对外提供的 gprc 服务端口
LOG_LEVEL=debug

.PHONY: build 
build:
	# 静态编译
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -o ${SRV} main.go

.PHONY: docker
docker:
	docker build . -t ${SRV}:${VERSION}

	docker tag ${SRV}:${VERSION} demo/${SRV}:${VERSION}
	docker push demo/${SRV}:${VERSION}

start:
	rm -rf logs
	dapr run --app-id ${SRV} \
         --app-protocol grpc \
         --app-port ${APP_PORT} \
         --dapr-http-port ${DAPR_HTTP_PORT} \
		 --dapr-grpc-port ${DAPR_GRPC_PORT} \
         --log-level ${LOG_LEVEL} \
         --components-path ./config \
         go run main.go

