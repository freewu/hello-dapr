## 服务介绍
```
提供一个 GET /hello 接口,接受 name 参数, 如果 name 为空 name 默认为 world
```

## 启动服务
```bash
git clone https://github.com/freewu/hello-dapr.git
cd exmaples/hello-dapr/hello-svr
make start
```

## 访问接口
```bash
curl http://127.0.0.1:9101/v1.0/invoke/demo-hello-srv/method/hello
```