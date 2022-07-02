package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"

	"OldPackageTest/grpc_test/proto"
)

func main() {
	//-------------配置导出jaeger---------------
	url := "http://47.106.87.191:14268/api/traces"
	tp, err := getTracerProviderToJaeger(url)
	//-------------配置导出jaeger---------------

	//-----------优雅关闭traceProvider------------
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context){
		ctx, cancel = context.WithTimeout(ctx, time.Second*1)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil{
			log.Fatal(err)
		}
	}(ctx)
	//-----------优雅关闭traceProvider------------

	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "flare"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
