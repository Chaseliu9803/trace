package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	
	"go.opentelemetry.io/otel"
	`go.opentelemetry.io/otel/attribute`
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	`go.opentelemetry.io/otel/sdk/resource`
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	`go.opentelemetry.io/otel/semconv`
	
	svc `otel/demo1/svc`
)

// 初始化 OpenTelemetry
func initTracer() *sdktrace.TracerProvider {
	exporter, err := jaeger.NewRawExporter(
		jaeger.WithAgentEndpoint(func(options *jaeger.AgentEndpointOptions) {
			options.Host = "localhost"
			options.Port = "6831"
		}),
	)
	if err != nil {
		log.Fatalf("Error creating Jaeger exporter: %v", err)
	}
	
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.ServiceNameKey.String("demo_service"), // 服务名
		)),
	)
	otel.SetTracerProvider(tp)
	return tp
}

func main() {
	tp := initTracer()
	defer func() {
		if cerr := tp.Shutdown(context.Background()); cerr != nil {
			log.Fatalf("Error shutting down tracer provider: %v", cerr)
		}
	}()
	
	//启动http服务器
	http.HandleFunc("/demo", handleRequest)
	
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Error starting Service A server: %v", err)
		}
	}()
	
	//模拟请求
	SimulateRequest()
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	tracer := otel.Tracer("root")
	//开始创建root span
	ctx, span := tracer.Start(req.Context(), "span root")
	defer span.End()
	
	//可以在span上记录一些信息，例如日志、请求参数、sql语句等
	span.SetAttributes(
		attribute.String("some root service info", "This is the root service"),
	)
	
	//访问服务A
	svc.CallServiceA(ctx)
	
	//访问服务B
	svc.CallServiceB(ctx)
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Response from Service Root")
}

func SimulateRequest()  {
	req, err := http.NewRequest("GET", "http://localhost:8080/demo", nil)
	if err != nil {
		log.Fatalf("Creating request fail: %v", err)
	}
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("Response received from Root Service")
}