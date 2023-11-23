package svc

import (
	`context`
	`fmt`
	
	`go.opentelemetry.io/otel`
	`go.opentelemetry.io/otel/attribute`
)

// Service A
func CallServiceA(ctx context.Context) {
	tracer := otel.Tracer("service A")
	ctx, span := tracer.Start(ctx, "ServiceA")
	defer span.End()
	
	span.SetAttributes(
		attribute.String("service info", "this is A service"),
	)
	
	// Call Service C
	CallServiceC(ctx)
	
	fmt.Println("Service A")
}


func CallServiceB(ctx context.Context) {
	tracer := otel.Tracer("Service B")
	ctx, span := tracer.Start(ctx, "ServiceB")
	defer span.End()
	
	span.SetAttributes(
		attribute.String("service info", "this is B service"),
	)
	
	CallServiceD(ctx)
	
	fmt.Println("Service B")
}

func CallServiceC(ctx context.Context) {
	tracer := otel.Tracer("Service C")
	ctx, span := tracer.Start(ctx, "ServiceC")
	defer span.End()
	
	span.SetAttributes(
		attribute.String("service info", "this is C service"),
	)
	
	CallServiceE(ctx)
	
	fmt.Println("Service C")
}



func CallServiceD(ctx context.Context) {
	tracer := otel.Tracer("Service D")
	ctx, span := tracer.Start(ctx, "ServiceD")
	defer span.End()
	
	span.SetAttributes(
		attribute.String("service info", "this is D service"),
	)
	
	fmt.Println("Service D")
}

func CallServiceE(ctx context.Context) {
	tracer := otel.Tracer("Service E")
	ctx, span := tracer.Start(ctx, "ServiceE")
	defer span.End()
	
	span.SetAttributes(
		attribute.String("service info", "this is E service"),
	)
	
	fmt.Println("Service E")
}
