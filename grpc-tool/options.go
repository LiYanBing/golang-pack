package grpc_tool

import (
	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

type Options struct {
	ServiceName string
	Tracer      opentracing.Tracer
	Log         log.Logger
}
