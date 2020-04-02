package grpc_tool

import (
	"github.com/go-kit/kit/endpoint"
)

type Wrapper func(options *Options, method string) endpoint.Middleware

var endpointMiddleware []Wrapper

func RegisterEndpointMiddleware(wrapper ...Wrapper) {
	if wrapper == nil {
		return
	}

	endpointMiddleware = append(endpointMiddleware, wrapper...)
}

func EndpointWrapperChain(options *Options, method string) []endpoint.Middleware {
	middlewares := make([]endpoint.Middleware, 0, len(endpointMiddleware))

	for _, wrapper := range endpointMiddleware {
		middlewares = append(middlewares, wrapper(options, method))
	}

	return middlewares
}
