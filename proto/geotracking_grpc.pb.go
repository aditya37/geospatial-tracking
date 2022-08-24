// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GeotrackingClient is the client API for Geotracking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeotrackingClient interface {
	GetDeviceLogByDeviceId(ctx context.Context, in *RequestGetDeviceLogByDeviceId, opts ...grpc.CallOption) (*ResponseGetDeviceLogByDeviceId, error)
	GetGPSTracking(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Geotracking_GetGPSTrackingClient, error)
	GetDeviceCounter(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseGetDeviceCounter, error)
}

type geotrackingClient struct {
	cc grpc.ClientConnInterface
}

func NewGeotrackingClient(cc grpc.ClientConnInterface) GeotrackingClient {
	return &geotrackingClient{cc}
}

func (c *geotrackingClient) GetDeviceLogByDeviceId(ctx context.Context, in *RequestGetDeviceLogByDeviceId, opts ...grpc.CallOption) (*ResponseGetDeviceLogByDeviceId, error) {
	out := new(ResponseGetDeviceLogByDeviceId)
	err := c.cc.Invoke(ctx, "/proto.Geotracking/GetDeviceLogByDeviceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geotrackingClient) GetGPSTracking(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Geotracking_GetGPSTrackingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Geotracking_ServiceDesc.Streams[0], "/proto.Geotracking/GetGPSTracking", opts...)
	if err != nil {
		return nil, err
	}
	x := &geotrackingGetGPSTrackingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Geotracking_GetGPSTrackingClient interface {
	Recv() (*ResponseStreamGPSTracking, error)
	grpc.ClientStream
}

type geotrackingGetGPSTrackingClient struct {
	grpc.ClientStream
}

func (x *geotrackingGetGPSTrackingClient) Recv() (*ResponseStreamGPSTracking, error) {
	m := new(ResponseStreamGPSTracking)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *geotrackingClient) GetDeviceCounter(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseGetDeviceCounter, error) {
	out := new(ResponseGetDeviceCounter)
	err := c.cc.Invoke(ctx, "/proto.Geotracking/GetDeviceCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeotrackingServer is the server API for Geotracking service.
// All implementations must embed UnimplementedGeotrackingServer
// for forward compatibility
type GeotrackingServer interface {
	GetDeviceLogByDeviceId(context.Context, *RequestGetDeviceLogByDeviceId) (*ResponseGetDeviceLogByDeviceId, error)
	GetGPSTracking(*emptypb.Empty, Geotracking_GetGPSTrackingServer) error
	GetDeviceCounter(context.Context, *emptypb.Empty) (*ResponseGetDeviceCounter, error)
	mustEmbedUnimplementedGeotrackingServer()
}

// UnimplementedGeotrackingServer must be embedded to have forward compatible implementations.
type UnimplementedGeotrackingServer struct {
}

func (UnimplementedGeotrackingServer) GetDeviceLogByDeviceId(context.Context, *RequestGetDeviceLogByDeviceId) (*ResponseGetDeviceLogByDeviceId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceLogByDeviceId not implemented")
}
func (UnimplementedGeotrackingServer) GetGPSTracking(*emptypb.Empty, Geotracking_GetGPSTrackingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetGPSTracking not implemented")
}
func (UnimplementedGeotrackingServer) GetDeviceCounter(context.Context, *emptypb.Empty) (*ResponseGetDeviceCounter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceCounter not implemented")
}
func (UnimplementedGeotrackingServer) mustEmbedUnimplementedGeotrackingServer() {}

// UnsafeGeotrackingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeotrackingServer will
// result in compilation errors.
type UnsafeGeotrackingServer interface {
	mustEmbedUnimplementedGeotrackingServer()
}

func RegisterGeotrackingServer(s grpc.ServiceRegistrar, srv GeotrackingServer) {
	s.RegisterService(&Geotracking_ServiceDesc, srv)
}

func _Geotracking_GetDeviceLogByDeviceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetDeviceLogByDeviceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeotrackingServer).GetDeviceLogByDeviceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Geotracking/GetDeviceLogByDeviceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeotrackingServer).GetDeviceLogByDeviceId(ctx, req.(*RequestGetDeviceLogByDeviceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geotracking_GetGPSTracking_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeotrackingServer).GetGPSTracking(m, &geotrackingGetGPSTrackingServer{stream})
}

type Geotracking_GetGPSTrackingServer interface {
	Send(*ResponseStreamGPSTracking) error
	grpc.ServerStream
}

type geotrackingGetGPSTrackingServer struct {
	grpc.ServerStream
}

func (x *geotrackingGetGPSTrackingServer) Send(m *ResponseStreamGPSTracking) error {
	return x.ServerStream.SendMsg(m)
}

func _Geotracking_GetDeviceCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeotrackingServer).GetDeviceCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Geotracking/GetDeviceCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeotrackingServer).GetDeviceCounter(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Geotracking_ServiceDesc is the grpc.ServiceDesc for Geotracking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geotracking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Geotracking",
	HandlerType: (*GeotrackingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceLogByDeviceId",
			Handler:    _Geotracking_GetDeviceLogByDeviceId_Handler,
		},
		{
			MethodName: "GetDeviceCounter",
			Handler:    _Geotracking_GetDeviceCounter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetGPSTracking",
			Handler:       _Geotracking_GetGPSTracking_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "geotracking.proto",
}
