// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/backend/backend.proto

package backend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Metric_AddFilm_FullMethodName  = "/backend.Metric/AddFilm"
	Metric_DelFilm_FullMethodName  = "/backend.Metric/DelFilm"
	Metric_GetFilms_FullMethodName = "/backend.Metric/GetFilms"
)

// MetricClient is the client API for Metric service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricClient interface {
	AddFilm(ctx context.Context, in *AddFilmRequest, opts ...grpc.CallOption) (*AddFilmResponse, error)
	DelFilm(ctx context.Context, in *DelFilmRequest, opts ...grpc.CallOption) (*DelFilmResponse, error)
	GetFilms(ctx context.Context, in *GetFilmsRequest, opts ...grpc.CallOption) (*GetFilmsResponse, error)
}

type metricClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricClient(cc grpc.ClientConnInterface) MetricClient {
	return &metricClient{cc}
}

func (c *metricClient) AddFilm(ctx context.Context, in *AddFilmRequest, opts ...grpc.CallOption) (*AddFilmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddFilmResponse)
	err := c.cc.Invoke(ctx, Metric_AddFilm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) DelFilm(ctx context.Context, in *DelFilmRequest, opts ...grpc.CallOption) (*DelFilmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelFilmResponse)
	err := c.cc.Invoke(ctx, Metric_DelFilm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) GetFilms(ctx context.Context, in *GetFilmsRequest, opts ...grpc.CallOption) (*GetFilmsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFilmsResponse)
	err := c.cc.Invoke(ctx, Metric_GetFilms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricServer is the server API for Metric service.
// All implementations must embed UnimplementedMetricServer
// for forward compatibility.
type MetricServer interface {
	AddFilm(context.Context, *AddFilmRequest) (*AddFilmResponse, error)
	DelFilm(context.Context, *DelFilmRequest) (*DelFilmResponse, error)
	GetFilms(context.Context, *GetFilmsRequest) (*GetFilmsResponse, error)
	mustEmbedUnimplementedMetricServer()
}

// UnimplementedMetricServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricServer struct{}

func (UnimplementedMetricServer) AddFilm(context.Context, *AddFilmRequest) (*AddFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFilm not implemented")
}
func (UnimplementedMetricServer) DelFilm(context.Context, *DelFilmRequest) (*DelFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFilm not implemented")
}
func (UnimplementedMetricServer) GetFilms(context.Context, *GetFilmsRequest) (*GetFilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilms not implemented")
}
func (UnimplementedMetricServer) mustEmbedUnimplementedMetricServer() {}
func (UnimplementedMetricServer) testEmbeddedByValue()                {}

// UnsafeMetricServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricServer will
// result in compilation errors.
type UnsafeMetricServer interface {
	mustEmbedUnimplementedMetricServer()
}

func RegisterMetricServer(s grpc.ServiceRegistrar, srv MetricServer) {
	// If the following call pancis, it indicates UnimplementedMetricServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Metric_ServiceDesc, srv)
}

func _Metric_AddFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).AddFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_AddFilm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).AddFilm(ctx, req.(*AddFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_DelFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).DelFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_DelFilm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).DelFilm(ctx, req.(*DelFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_GetFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).GetFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_GetFilms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).GetFilms(ctx, req.(*GetFilmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metric_ServiceDesc is the grpc.ServiceDesc for Metric service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metric_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Metric",
	HandlerType: (*MetricServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFilm",
			Handler:    _Metric_AddFilm_Handler,
		},
		{
			MethodName: "DelFilm",
			Handler:    _Metric_DelFilm_Handler,
		},
		{
			MethodName: "GetFilms",
			Handler:    _Metric_GetFilms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/backend/backend.proto",
}
