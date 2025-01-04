// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api/content/user_content_service.proto

package content

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserContentService_GetList_FullMethodName   = "/content.UserContentService/GetList"
	UserContentService_GetValued_FullMethodName = "/content.UserContentService/GetValued"
	UserContentService_Add_FullMethodName       = "/content.UserContentService/Add"
	UserContentService_Remove_FullMethodName    = "/content.UserContentService/Remove"
)

// UserContentServiceClient is the client API for UserContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserContentServiceClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetValued(ctx context.Context, in *GetValuedRequest, opts ...grpc.CallOption) (*GetValuedResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Remove(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userContentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserContentServiceClient(cc grpc.ClientConnInterface) UserContentServiceClient {
	return &userContentServiceClient{cc}
}

func (c *userContentServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, UserContentService_GetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userContentServiceClient) GetValued(ctx context.Context, in *GetValuedRequest, opts ...grpc.CallOption) (*GetValuedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetValuedResponse)
	err := c.cc.Invoke(ctx, UserContentService_GetValued_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userContentServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserContentService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userContentServiceClient) Remove(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserContentService_Remove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserContentServiceServer is the server API for UserContentService service.
// All implementations must embed UnimplementedUserContentServiceServer
// for forward compatibility.
type UserContentServiceServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetValued(context.Context, *GetValuedRequest) (*GetValuedResponse, error)
	Add(context.Context, *AddRequest) (*emptypb.Empty, error)
	Remove(context.Context, *RemoveItemRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserContentServiceServer()
}

// UnimplementedUserContentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserContentServiceServer struct{}

func (UnimplementedUserContentServiceServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedUserContentServiceServer) GetValued(context.Context, *GetValuedRequest) (*GetValuedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValued not implemented")
}
func (UnimplementedUserContentServiceServer) Add(context.Context, *AddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedUserContentServiceServer) Remove(context.Context, *RemoveItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedUserContentServiceServer) mustEmbedUnimplementedUserContentServiceServer() {}
func (UnimplementedUserContentServiceServer) testEmbeddedByValue()                            {}

// UnsafeUserContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserContentServiceServer will
// result in compilation errors.
type UnsafeUserContentServiceServer interface {
	mustEmbedUnimplementedUserContentServiceServer()
}

func RegisterUserContentServiceServer(s grpc.ServiceRegistrar, srv UserContentServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserContentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserContentService_ServiceDesc, srv)
}

func _UserContentService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserContentServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserContentService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserContentServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserContentService_GetValued_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValuedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserContentServiceServer).GetValued(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserContentService_GetValued_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserContentServiceServer).GetValued(ctx, req.(*GetValuedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserContentService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserContentServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserContentService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserContentServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserContentService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserContentServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserContentService_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserContentServiceServer).Remove(ctx, req.(*RemoveItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserContentService_ServiceDesc is the grpc.ServiceDesc for UserContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.UserContentService",
	HandlerType: (*UserContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _UserContentService_GetList_Handler,
		},
		{
			MethodName: "GetValued",
			Handler:    _UserContentService_GetValued_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _UserContentService_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _UserContentService_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/content/user_content_service.proto",
}
