package main

import (
 "context"
 "log"
 "time"

 "go.opentelemetry.io/otel"
 "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
 "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
 "go.opentelemetry.io/otel/sdk/resource"
 "go.opentelemetry.io/otel/sdk/trace"
 semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func startTracing() (*trace.TracerProvider, error) {
 headers := map[string]string{
  "content-type": "application/json",
 }

 exporter, err := otlptrace.New(
  context.Background(),
  otlptracehttp.NewClient(
   otlptracehttp.WithEndpoint("localhost:4318"),
   otlptracehttp.WithHeaders(headers),
   otlptracehttp.WithInsecure(),
  ),
 )
 if err != nil {
  log.Panic("Error creating new exporter: %w", err)
 }

 tracerprovider := trace.NewTracerProvider(
  trace.WithBatcher(
   exporter,
   trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
   trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
   trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
  ),
  trace.WithResource(
   resource.NewWithAttributes(
    semconv.SchemaURL,
    semconv.ServiceNameKey.String("test-app"),
   ),
  ),
 )

 otel.SetTracerProvider(tracerprovider)

 return tracerprovider, nil
}

func main(){

	log.Println("running pusher client..")

	tp, err := startTracing();

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Panic(err)
		}
	}()

	tracer := tp.Tracer("pusher-jaeger")

 	_, span := tracer.Start(context.Background(), "Test Span from Pusher")
	defer span.End()

	log.Println("sending spans..")
	log.Println("done!")
}
