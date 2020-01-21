// GRPC网关grpc-gateway是protoc的一个插件。它遵循gRPC中的服务定义，生成反向代理服务，这个代理就会把RESTful风格的JSON API转成gRPC请求。

// 我们使用go-grpc写后台服务。Go-GRPC是客户端与服务端的go-micro、grpc插件包装器。当调用grpc.NewService时，它会返回micro.Service服务。

// 在下面的代码中我们写了一个简单的问候（greeter）服务API。类似的代码也会用来注册其它的端点。切记，网关需要端点地址及服务。

package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
    hello "github.com/micro/examples/grpc/gateway/proto/hello"
    // hello "grpc-gateway/routes/grpc-gateway"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
)

var (
	endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
