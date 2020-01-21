package main


/* 项目启动

1、运行greeter服务
go run greeter.go --registry=mdns --server_address=localhost:9090


2、运行网关程序，把greeter服务所有端点都放在localhost:9090下
go run main.go

*/


/* 测试请求

请求参数json格式
curl -d '{"name": "john"}' http://localhost:8080/greeter/hello

*/