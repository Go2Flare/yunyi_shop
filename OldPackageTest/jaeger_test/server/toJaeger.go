package main

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"log"
)

const (
	Service     = "otel-trace-server-demo"
	environment = "test"
	id          = 1
)

//上传多个jaeger endpoint
func getTracerProviderToJaeger(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	log.Printf("upload trace to jaeger : %v \n", url)
	tp := tracesdk.NewTracerProvider(
		// 使用jaeger的exporter
		tracesdk.WithBatcher(exp),
		// 配置应用相关信息，指标来源
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(Service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)

	//最后得设置全局的provider
	otel.SetTracerProvider(tp)
	return tp, nil
}