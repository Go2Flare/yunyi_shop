package main

import (
	"OldPackageTest/grpc_test/proto"
	"context"
	"errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

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

	_, span := otel.Tracer("server").Start(ctx, "grpc-server")
	defer span.End()

	/*----------------grpc服务---------------*/
	g := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
		)

	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		span.RecordError(errors.New("net.Listen err"), trace.WithTimestamp(time.Now()))
		span.SetStatus(codes.Error, err.Error())
		panic("failed to listen:" + err.Error())
	}else{
		//span.AddEvent("log", trace.WithAttributes(
		//	attribute.String("log.severity", "info"),
		//	attribute.String("log.message", "grpc serving successfully"),
		//))
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
	/*----------------grpc服务---------------*/
}
