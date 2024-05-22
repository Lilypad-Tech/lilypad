package lilymetrics

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	// "go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

// var (
// 	tracer       trace.Tracer
// 	otlpEndpoint string
// )

type BasicAuthTransport struct {
	Username string
	Password string
	Base     http.RoundTripper
}

func newConsoleExporter() (oteltrace.SpanExporter, error) {
	return stdouttrace.New()
	// return nil, nil
}

//	func GrafanaExporter() (oteltrace.SpanExporter, error) {
//		exporter, err := jaeger.NewRawExporter(
//			jaeger.WithCollectorEndpoint("http://jaeger-collector:14268/api/traces"), // Adjust endpoint as needed
//			jaeger.WithUsername("your_username_here"),                                // Add your username
//			jaeger.WithPassword("your_password_here"),                                // Add your password
//			jaeger.WithHTTPClient(&http.Client{
//				Transport: &http.Transport{
//					Proxy: http.ProxyFromEnvironment,
//				},
//			}),
//		)
//		return exporter, err
//	}
func newOTLPExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	// Change default HTTPS -> HTTP
	insecureOpt := otlptracehttp.WithInsecure()
	// auth := otlptracehttp.NewClient()
	// Update default OTLP reciver endpoint
	//
	endpointOpt := otlptracehttp.WithEndpoint(os.Getenv("TEMPO") + ":4318")

	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)

}
func (bat *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(bat.Username, bat.Password)
	return bat.Base.RoundTrip(req)
}
func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(os.Args[1]), //"Metrics"),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}
func getCallStack() []string {
	// Set the maximum depth of the call stack to retrieve
	const depth = 32

	// Create a slice to store the function names
	stack := make([]uintptr, depth)

	// Retrieve the call stack
	n := runtime.Callers(0, stack)
	if n == 0 {
		return nil
	}

	// Trim the stack to the actual size
	stack = stack[:n]

	// Retrieve function names for each entry in the call stack
	var functionNames []string
	for _, pc := range stack {
		if fn := runtime.FuncForPC(pc); fn != nil {
			functionNames = append(functionNames, fn.Name())
		}
	}

	return functionNames

}
func Trace(ctx context.Context) trace.Span {
	return TraceSection(ctx, "")
}
func TraceSection(ctx context.Context, name string) trace.Span {

	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	if path.Base(funcName) == "lilymetrics.Trace" {
		pc, _, _, _ = runtime.Caller(2)
		funcName = runtime.FuncForPC(pc).Name()
	}
	// tracer := otel.Tracer(strings.Join(getCallStack(), "->"))

	// Example context, replace with your context creation method
	// ctx := context.Background()
	exp, _ := newOTLPExporter(context.Background()) //newConsoleExporter()
	tp := newTraceProvider(exp)
	otel.SetTracerProvider(tp)

	tracer := tp.Tracer(funcName)

	_, span := tracer.Start(ctx, path.Base(funcName)+" "+name)
	// fmt.Println(p, name)
	return span
}
func LogDeal(source string, deal, status string) {
	Post("metrics-dashboard/deal", fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, deal, status))
}

func Matcher(jobOffers int, resourceOffers int, deals int) {
	Post("metrics-dashboard/matcher", fmt.Sprintf(`{"jobOffers":"%d","resourceOffers":"%d","deals":"%d"}`, jobOffers, resourceOffers, deals))
}
func LogJob(dealid string, state string, metadata string) {
	Post("metrics-dashboard/job", fmt.Sprintf(`{"dealid":"%s","state":"%s","metadata":"%s"}`, dealid, state, metadata))
}
func LogJobStatus(job_id string, status string, module_id string) {
	Post("metrics-dashboard/status", fmt.Sprintf(`{"job_id":"%s","status":"%s","module_id":"%s"}`, job_id, status, module_id))
}
func LogMetric(module_id string, detail string) {
	Post("metrics-dashboard/log", fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, module_id, detail))
}
func LogResult(module_id string, detail string) {
	Post("metrics-dashboard/result", fmt.Sprintf(`{"Type":"%s","result":%s}`, module_id, detail))
}
func Post(path string, json string) {
	data := []byte(json)
	url := os.Getenv("METRICS_HOST") + ":8000/" + path
	client := &http.Client{Timeout: time.Second * 1}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
	} else {
		resp.Body.Close()
	}
}
