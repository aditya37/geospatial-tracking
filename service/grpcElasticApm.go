package service

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

func GetGrpcServerElasticApmOptions(
	interceptor []grpc.UnaryServerInterceptor,
	interceptors []grpc.StreamServerInterceptor,
) []grpc.ServerOption {
	return []grpc.ServerOption{
		// middleware for grpc unary request
		grpc_middleware.WithUnaryServerChain(
			getDefaultUnaryOption(interceptor...)...,
		),
		grpc_middleware.WithStreamServerChain(
			getDefaultStreamOption(interceptors...)...,
		),
	}
}

//SetUnaryMiddleware
// set or add middleware for unary request....
func SetUnaryMiddleware(interceptor ...grpc.UnaryServerInterceptor) []grpc.UnaryServerInterceptor {
	interceptor = append(interceptor, interceptor...)
	return interceptor
}

// SetStreamMiddleware
func SetStreamMiddleware(interceptors ...grpc.StreamServerInterceptor) []grpc.StreamServerInterceptor {
	interceptors = append(interceptors, interceptors...)
	return interceptors
}

// getDefaultStreamOption
// default option for stream interceptor
func getDefaultStreamOption(interceptors ...grpc.StreamServerInterceptor) []grpc.StreamServerInterceptor {
	interceptors = append(
		interceptors,
		apmgrpc.NewStreamServerInterceptor(apmgrpc.WithRecovery()),
	)
	return interceptors
}

// getDefaultUnaryOption...
// default option for unary interceptor
func getDefaultUnaryOption(interceptor ...grpc.UnaryServerInterceptor) []grpc.UnaryServerInterceptor {
	interceptor = append(
		interceptor,
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
	)

	return interceptor
}
