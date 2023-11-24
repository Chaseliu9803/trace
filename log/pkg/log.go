package pkg

import (
	`context`
	`log`
	`os`
	
	`go.opentelemetry.io/otel/attribute`
	`go.opentelemetry.io/otel/trace`
)

//封装一下日志
type Logger struct {
	log *log.Logger
}

var Log *Logger

func init()  {
	Log = &Logger{
		log: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("", trace.WithAttributes(attribute.String("log", msg)))
	l.log.Println(msg)
}

//todo
func (l *Logger) Info(ctx context.Context, msg string) {}
//todo
func (l *Logger) Warn(ctx context.Context, msg string) {}
//todo
func (l *Logger) Error(ctx context.Context, msg string) {}
//todo
func (l *Logger) Fatal(ctx context.Context, msg string) {}


