package lilymetrics

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	// "go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
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
	endpointOpt := otlptracehttp.WithEndpoint("localhost:4318")

	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
	// otlpEndpoint = os.Getenv("OTLP_ENDPOINT")
	// if otlpEndpoint == "" {
	// 	log.Fatalln("You MUST set OTLP_ENDPOINT env variable!")
	// }
	// username := os.Getenv("OTLP_USERNAME")
	// password := os.Getenv("OTLP_PASSWORD")
	// // Create an HTTP client with basic authentication
	// httpClient := &http.Client{
	// 	Transport: &http.Transport{
	// 		Proxy: http.ProxyFromEnvironment,
	// 	},
	// }

	// // Set up basic authentication
	// httpClient.Transport = &BasicAuthTransport{
	// 	Username: username,
	// 	Password: password,
	// 	Base:     httpClient.Transport,
	// }

	// //Initialize the exporter with the HTTP client
	// // otlptracehttp.NewClient()

	// auth := username + ":" + password
	// authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	// authHeader := "Basic " + authEncoded

	// exporter, err := otlptracehttp.New(
	// 	ctx,
	// 	otlptracehttp.WithHeaders(map[string]string{
	// 		"Authorization": authHeader,
	// 	}),
	// 	otlptracehttp.WithEndpoint(otlpEndpoint),
	// )

	// if err != nil {
	// 	// handle error
	// }
	// return exporter, err
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
	log.Print(deal)
	url := "http://" + os.Getenv("METRICS_HOST") + ":8000/metrics-dashboard/deal"
	json := fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, deal, status)
	fmt.Println(json)
	data := []byte(json)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func LogJob(source string, module_id string, job_id string, status string) {
	log.Print(module_id)

	url := "http://" + os.Getenv("METRICS_HOST") + ":8000/metrics-dashboard/job"
	json := fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, module_id, status)
	fmt.Println(json)
	data := []byte(json)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func LogMetric(module_id string, detail string) {
	log.Print(module_id)
	url := "http://" + os.Getenv("METRICS_HOST") + ":8000/metrics-dashboard/log"
	json := fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, module_id, detail)
	fmt.Println(json)
	data := []byte(json)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
