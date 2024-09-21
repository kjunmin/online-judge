// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: common/api/judge.proto

package api

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
	ProblemsService_CreateProblem_FullMethodName   = "/api.ProblemsService/CreateProblem"
	ProblemsService_GetProblemsList_FullMethodName = "/api.ProblemsService/GetProblemsList"
	ProblemsService_GetProblemById_FullMethodName  = "/api.ProblemsService/GetProblemById"
)

// ProblemsServiceClient is the client API for ProblemsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemsServiceClient interface {
	CreateProblem(ctx context.Context, in *CreateProblemRequest, opts ...grpc.CallOption) (*Problem, error)
	GetProblemsList(ctx context.Context, in *GetProblemsListRequest, opts ...grpc.CallOption) (*GetProblemsListResponse, error)
	GetProblemById(ctx context.Context, in *GetProblemByIdRequest, opts ...grpc.CallOption) (*Problem, error)
}

type problemsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemsServiceClient(cc grpc.ClientConnInterface) ProblemsServiceClient {
	return &problemsServiceClient{cc}
}

func (c *problemsServiceClient) CreateProblem(ctx context.Context, in *CreateProblemRequest, opts ...grpc.CallOption) (*Problem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Problem)
	err := c.cc.Invoke(ctx, ProblemsService_CreateProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemsServiceClient) GetProblemsList(ctx context.Context, in *GetProblemsListRequest, opts ...grpc.CallOption) (*GetProblemsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProblemsListResponse)
	err := c.cc.Invoke(ctx, ProblemsService_GetProblemsList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemsServiceClient) GetProblemById(ctx context.Context, in *GetProblemByIdRequest, opts ...grpc.CallOption) (*Problem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Problem)
	err := c.cc.Invoke(ctx, ProblemsService_GetProblemById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemsServiceServer is the server API for ProblemsService service.
// All implementations must embed UnimplementedProblemsServiceServer
// for forward compatibility.
type ProblemsServiceServer interface {
	CreateProblem(context.Context, *CreateProblemRequest) (*Problem, error)
	GetProblemsList(context.Context, *GetProblemsListRequest) (*GetProblemsListResponse, error)
	GetProblemById(context.Context, *GetProblemByIdRequest) (*Problem, error)
	mustEmbedUnimplementedProblemsServiceServer()
}

// UnimplementedProblemsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProblemsServiceServer struct{}

func (UnimplementedProblemsServiceServer) CreateProblem(context.Context, *CreateProblemRequest) (*Problem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProblem not implemented")
}
func (UnimplementedProblemsServiceServer) GetProblemsList(context.Context, *GetProblemsListRequest) (*GetProblemsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemsList not implemented")
}
func (UnimplementedProblemsServiceServer) GetProblemById(context.Context, *GetProblemByIdRequest) (*Problem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemById not implemented")
}
func (UnimplementedProblemsServiceServer) mustEmbedUnimplementedProblemsServiceServer() {}
func (UnimplementedProblemsServiceServer) testEmbeddedByValue()                         {}

// UnsafeProblemsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemsServiceServer will
// result in compilation errors.
type UnsafeProblemsServiceServer interface {
	mustEmbedUnimplementedProblemsServiceServer()
}

func RegisterProblemsServiceServer(s grpc.ServiceRegistrar, srv ProblemsServiceServer) {
	// If the following call pancis, it indicates UnimplementedProblemsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProblemsService_ServiceDesc, srv)
}

func _ProblemsService_CreateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemsServiceServer).CreateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemsService_CreateProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemsServiceServer).CreateProblem(ctx, req.(*CreateProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemsService_GetProblemsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemsServiceServer).GetProblemsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemsService_GetProblemsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemsServiceServer).GetProblemsList(ctx, req.(*GetProblemsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemsService_GetProblemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemsServiceServer).GetProblemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemsService_GetProblemById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemsServiceServer).GetProblemById(ctx, req.(*GetProblemByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemsService_ServiceDesc is the grpc.ServiceDesc for ProblemsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProblemsService",
	HandlerType: (*ProblemsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProblem",
			Handler:    _ProblemsService_CreateProblem_Handler,
		},
		{
			MethodName: "GetProblemsList",
			Handler:    _ProblemsService_GetProblemsList_Handler,
		},
		{
			MethodName: "GetProblemById",
			Handler:    _ProblemsService_GetProblemById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/api/judge.proto",
}
