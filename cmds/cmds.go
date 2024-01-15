package cmds

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"gitlab.com/kulyklev/autoria-parser/cmds/migrate"
	"gitlab.com/kulyklev/autoria-parser/cmds/parse"
	"gitlab.com/kulyklev/autoria-parser/cmds/parse_failed"

	"github.com/urfave/cli/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.14.0"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

func Exec(args []string, log *zap.SugaredLogger) error {
	// =================================================================================================================
	// Start Tracing Support
	log.Infow("startup", "status", "initializing OT/Jaeger tracing support")

	traceProvider, err := startTracing(
		log,
		"parser-cli",
		"http://localhost:14268/api/traces",
		1,
	)
	if err != nil {
		return fmt.Errorf("starting tracing: %w", err)
	}
	defer traceProvider.Shutdown(context.Background())

	tracer := traceProvider.Tracer("cli-parser")

	// =================================================================================================================
	// GOMAXPROCS
	opt := maxprocs.Logger(log.Infof)
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(runtime.NumCPU()))

	app := &cli.App{
		Name:        "auto-ria parser",
		Usage:       "command-name admin tools",
		UsageText:   "command [options] [arguments]",
		Description: "Set of commands to parse auto-ria car prices for specified search",
		Commands: []*cli.Command{
			migrate.Commands(),
			parse.CommandParse(log, tracer),
			parse_failed.CommandParse(log),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "url",
				Usage: "connection url for SQLServer",
				EnvVars: []string{
					"SQL_URL",
				},
			},
		},
	}

	return app.Run(args)
}

// =============================================================================+=======================================

// startTracing configure open telemetry to be used with tracing provider.
func startTracing(log *zap.SugaredLogger, serviceName string, reporterURI string, probability float64) (*trace.TracerProvider, error) {
	//withLogger := jaeger.WithLogger(zap.NewStdLog(log))
	endpoint := jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(reporterURI))
	exporter, err := jaeger.New(endpoint)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(probability)),
		trace.WithBatcher(exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				attribute.String("exporter", "jaeger"),
			),
		),
	)

	otel.SetTracerProvider(traceProvider)
	return traceProvider, nil
}
