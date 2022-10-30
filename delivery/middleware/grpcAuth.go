package middleware

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// contenxt key....
type MiddlewareCtxKey string

var (
	CtxKeyStreamerId MiddlewareCtxKey = "streamer_id"
)

func (mk MiddlewareCtxKey) ToString() string {
	return string(mk)
}

// error variable...
var (
	ErrHeaderStreamerIdNotFound = errors.New("streamer id not found")
	ErrFailedGetHeader          = errors.New("failed get header")
)

// grpc stream custom context
type customStreamContenxt struct {
	// assign interface grpc stream
	grpc.ServerStream
	WrapContext context.Context
}

// implement method Context from interface grpc.ServerStream
func (cc *customStreamContenxt) Context() context.Context {
	return cc.WrapContext
}
func wrapServerStreamContext(stream grpc.ServerStream) *customStreamContenxt {
	// assign struct with interface grpc stream
	if exists, ok := stream.(*customStreamContenxt); ok {
		return exists
	}
	return &customStreamContenxt{
		ServerStream: stream,
		WrapContext:  stream.Context(),
	}
}

func UnaryAuthMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		return handler(ctx, req)
	}
}

func StreamAuthMiddleware() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		streamCtx := ss.Context()
		var newCtx context.Context
		streamerId, err := getHeaderStreamerId(streamCtx)
		if err != nil {
			return ErrFailedGetHeader
		}

		// implement custom stream context
		customCtx := wrapServerStreamContext(ss)

		newCtx = context.WithValue(
			streamCtx,
			CtxKeyStreamerId.ToString(),
			streamerId,
		)
		customCtx.WrapContext = newCtx

		return handler(srv, customCtx)
	}
}

func getHeaderStreamerId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrFailedGetHeader
	}
	header := md["streamer_id"]
	if len(header) <= 0 || header[0] == "" {
		return "", ErrHeaderStreamerIdNotFound
	}
	return header[0], nil
}
