package grpc_tool

import (
	"github.com/go-kit/kit/transport/grpc"
)

type ClientRequestFunc func(options *Options, method string) grpc.ClientRequestFunc

type ServerRequestFunc func(options *Options, method string) grpc.ServerRequestFunc

var (
	clientRequestFunc []ClientRequestFunc
	serverRequestFunc []ServerRequestFunc
)

func RegisterClientRequestFunc(fun ...ClientRequestFunc) {
	if fun == nil {
		return
	}

	clientRequestFunc = append(clientRequestFunc, fun...)
}

func RegisterServerRequestFunc(fun ...ServerRequestFunc) {
	if fun == nil {
		return
	}

	serverRequestFunc = append(serverRequestFunc, fun...)
}

func ContextToGRPC(options *Options, method string) []grpc.ClientRequestFunc {
	funcs := make([]grpc.ClientRequestFunc, 0, len(clientRequestFunc))

	for _, mid := range clientRequestFunc {
		funcs = append(funcs, mid(options, method))
	}
	return funcs
}

func GRPCToContext(options *Options, method string) []grpc.ServerRequestFunc {
	funcs := make([]grpc.ServerRequestFunc, 0, len(serverRequestFunc))

	for _, mid := range serverRequestFunc {
		funcs = append(funcs, mid(options, method))
	}

	return funcs
}
