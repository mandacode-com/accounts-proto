// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: auth/v1/user.proto

package authv1

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
	LocalUserService_CreateLocalUser_FullMethodName         = "/auth.v1.LocalUserService/CreateLocalUser"
	LocalUserService_DeleteLocalUser_FullMethodName         = "/auth.v1.LocalUserService/DeleteLocalUser"
	LocalUserService_UpdateLocalUserEmail_FullMethodName    = "/auth.v1.LocalUserService/UpdateLocalUserEmail"
	LocalUserService_UpdateEmailVerification_FullMethodName = "/auth.v1.LocalUserService/UpdateEmailVerification"
)

// LocalUserServiceClient is the client API for LocalUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalUserServiceClient interface {
	// CreateLocalUser creates a new local user with the provided details
	CreateLocalUser(ctx context.Context, in *CreateLocalUserRequest, opts ...grpc.CallOption) (*CreateLocalUserResponse, error)
	// DeleteLocalUser deletes a local user by user ID
	DeleteLocalUser(ctx context.Context, in *DeleteLocalUserRequest, opts ...grpc.CallOption) (*DeleteLocalUserResponse, error)
	// UpdateLocalUserEmail updates the user's email address
	UpdateLocalUserEmail(ctx context.Context, in *UpdateLocalUserEmailRequest, opts ...grpc.CallOption) (*UpdateLocalUserEmailResponse, error)
	// UpdateEmailVerification sets the email verification status for a user
	UpdateEmailVerification(ctx context.Context, in *UpdateEmailVerificationRequest, opts ...grpc.CallOption) (*UpdateEmailVerificationResponse, error)
}

type localUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalUserServiceClient(cc grpc.ClientConnInterface) LocalUserServiceClient {
	return &localUserServiceClient{cc}
}

func (c *localUserServiceClient) CreateLocalUser(ctx context.Context, in *CreateLocalUserRequest, opts ...grpc.CallOption) (*CreateLocalUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLocalUserResponse)
	err := c.cc.Invoke(ctx, LocalUserService_CreateLocalUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localUserServiceClient) DeleteLocalUser(ctx context.Context, in *DeleteLocalUserRequest, opts ...grpc.CallOption) (*DeleteLocalUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLocalUserResponse)
	err := c.cc.Invoke(ctx, LocalUserService_DeleteLocalUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localUserServiceClient) UpdateLocalUserEmail(ctx context.Context, in *UpdateLocalUserEmailRequest, opts ...grpc.CallOption) (*UpdateLocalUserEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLocalUserEmailResponse)
	err := c.cc.Invoke(ctx, LocalUserService_UpdateLocalUserEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localUserServiceClient) UpdateEmailVerification(ctx context.Context, in *UpdateEmailVerificationRequest, opts ...grpc.CallOption) (*UpdateEmailVerificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEmailVerificationResponse)
	err := c.cc.Invoke(ctx, LocalUserService_UpdateEmailVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalUserServiceServer is the server API for LocalUserService service.
// All implementations must embed UnimplementedLocalUserServiceServer
// for forward compatibility.
type LocalUserServiceServer interface {
	// CreateLocalUser creates a new local user with the provided details
	CreateLocalUser(context.Context, *CreateLocalUserRequest) (*CreateLocalUserResponse, error)
	// DeleteLocalUser deletes a local user by user ID
	DeleteLocalUser(context.Context, *DeleteLocalUserRequest) (*DeleteLocalUserResponse, error)
	// UpdateLocalUserEmail updates the user's email address
	UpdateLocalUserEmail(context.Context, *UpdateLocalUserEmailRequest) (*UpdateLocalUserEmailResponse, error)
	// UpdateEmailVerification sets the email verification status for a user
	UpdateEmailVerification(context.Context, *UpdateEmailVerificationRequest) (*UpdateEmailVerificationResponse, error)
	mustEmbedUnimplementedLocalUserServiceServer()
}

// UnimplementedLocalUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocalUserServiceServer struct{}

func (UnimplementedLocalUserServiceServer) CreateLocalUser(context.Context, *CreateLocalUserRequest) (*CreateLocalUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocalUser not implemented")
}
func (UnimplementedLocalUserServiceServer) DeleteLocalUser(context.Context, *DeleteLocalUserRequest) (*DeleteLocalUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocalUser not implemented")
}
func (UnimplementedLocalUserServiceServer) UpdateLocalUserEmail(context.Context, *UpdateLocalUserEmailRequest) (*UpdateLocalUserEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocalUserEmail not implemented")
}
func (UnimplementedLocalUserServiceServer) UpdateEmailVerification(context.Context, *UpdateEmailVerificationRequest) (*UpdateEmailVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailVerification not implemented")
}
func (UnimplementedLocalUserServiceServer) mustEmbedUnimplementedLocalUserServiceServer() {}
func (UnimplementedLocalUserServiceServer) testEmbeddedByValue()                          {}

// UnsafeLocalUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalUserServiceServer will
// result in compilation errors.
type UnsafeLocalUserServiceServer interface {
	mustEmbedUnimplementedLocalUserServiceServer()
}

func RegisterLocalUserServiceServer(s grpc.ServiceRegistrar, srv LocalUserServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocalUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocalUserService_ServiceDesc, srv)
}

func _LocalUserService_CreateLocalUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocalUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalUserServiceServer).CreateLocalUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalUserService_CreateLocalUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalUserServiceServer).CreateLocalUser(ctx, req.(*CreateLocalUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalUserService_DeleteLocalUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocalUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalUserServiceServer).DeleteLocalUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalUserService_DeleteLocalUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalUserServiceServer).DeleteLocalUser(ctx, req.(*DeleteLocalUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalUserService_UpdateLocalUserEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocalUserEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalUserServiceServer).UpdateLocalUserEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalUserService_UpdateLocalUserEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalUserServiceServer).UpdateLocalUserEmail(ctx, req.(*UpdateLocalUserEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalUserService_UpdateEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalUserServiceServer).UpdateEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalUserService_UpdateEmailVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalUserServiceServer).UpdateEmailVerification(ctx, req.(*UpdateEmailVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalUserService_ServiceDesc is the grpc.ServiceDesc for LocalUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.LocalUserService",
	HandlerType: (*LocalUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocalUser",
			Handler:    _LocalUserService_CreateLocalUser_Handler,
		},
		{
			MethodName: "DeleteLocalUser",
			Handler:    _LocalUserService_DeleteLocalUser_Handler,
		},
		{
			MethodName: "UpdateLocalUserEmail",
			Handler:    _LocalUserService_UpdateLocalUserEmail_Handler,
		},
		{
			MethodName: "UpdateEmailVerification",
			Handler:    _LocalUserService_UpdateEmailVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/user.proto",
}

const (
	OAuthUserService_CreateOAuthUser_FullMethodName = "/auth.v1.OAuthUserService/CreateOAuthUser"
	OAuthUserService_DeleteOAuthUser_FullMethodName = "/auth.v1.OAuthUserService/DeleteOAuthUser"
	OAuthUserService_SyncOAuthUser_FullMethodName   = "/auth.v1.OAuthUserService/SyncOAuthUser"
)

// OAuthUserServiceClient is the client API for OAuthUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthUserServiceClient interface {
	// CreateOAuthUser creates a new OAuth user with the provided details
	CreateOAuthUser(ctx context.Context, in *CreateOAuthUserRequest, opts ...grpc.CallOption) (*CreateOAuthUserResponse, error)
	// DeleteOAuthUser deletes an OAuth user by user ID
	DeleteOAuthUser(ctx context.Context, in *DeleteOAuthUserRequest, opts ...grpc.CallOption) (*DeleteOAuthUserResponse, error)
	// SyncOAuthUser synchronizes an OAuth user with the provided details
	SyncOAuthUser(ctx context.Context, in *SyncOAuthUserRequest, opts ...grpc.CallOption) (*SyncOAuthUserResponse, error)
}

type oAuthUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthUserServiceClient(cc grpc.ClientConnInterface) OAuthUserServiceClient {
	return &oAuthUserServiceClient{cc}
}

func (c *oAuthUserServiceClient) CreateOAuthUser(ctx context.Context, in *CreateOAuthUserRequest, opts ...grpc.CallOption) (*CreateOAuthUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOAuthUserResponse)
	err := c.cc.Invoke(ctx, OAuthUserService_CreateOAuthUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthUserServiceClient) DeleteOAuthUser(ctx context.Context, in *DeleteOAuthUserRequest, opts ...grpc.CallOption) (*DeleteOAuthUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteOAuthUserResponse)
	err := c.cc.Invoke(ctx, OAuthUserService_DeleteOAuthUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthUserServiceClient) SyncOAuthUser(ctx context.Context, in *SyncOAuthUserRequest, opts ...grpc.CallOption) (*SyncOAuthUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncOAuthUserResponse)
	err := c.cc.Invoke(ctx, OAuthUserService_SyncOAuthUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthUserServiceServer is the server API for OAuthUserService service.
// All implementations must embed UnimplementedOAuthUserServiceServer
// for forward compatibility.
type OAuthUserServiceServer interface {
	// CreateOAuthUser creates a new OAuth user with the provided details
	CreateOAuthUser(context.Context, *CreateOAuthUserRequest) (*CreateOAuthUserResponse, error)
	// DeleteOAuthUser deletes an OAuth user by user ID
	DeleteOAuthUser(context.Context, *DeleteOAuthUserRequest) (*DeleteOAuthUserResponse, error)
	// SyncOAuthUser synchronizes an OAuth user with the provided details
	SyncOAuthUser(context.Context, *SyncOAuthUserRequest) (*SyncOAuthUserResponse, error)
	mustEmbedUnimplementedOAuthUserServiceServer()
}

// UnimplementedOAuthUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOAuthUserServiceServer struct{}

func (UnimplementedOAuthUserServiceServer) CreateOAuthUser(context.Context, *CreateOAuthUserRequest) (*CreateOAuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOAuthUser not implemented")
}
func (UnimplementedOAuthUserServiceServer) DeleteOAuthUser(context.Context, *DeleteOAuthUserRequest) (*DeleteOAuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOAuthUser not implemented")
}
func (UnimplementedOAuthUserServiceServer) SyncOAuthUser(context.Context, *SyncOAuthUserRequest) (*SyncOAuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncOAuthUser not implemented")
}
func (UnimplementedOAuthUserServiceServer) mustEmbedUnimplementedOAuthUserServiceServer() {}
func (UnimplementedOAuthUserServiceServer) testEmbeddedByValue()                          {}

// UnsafeOAuthUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthUserServiceServer will
// result in compilation errors.
type UnsafeOAuthUserServiceServer interface {
	mustEmbedUnimplementedOAuthUserServiceServer()
}

func RegisterOAuthUserServiceServer(s grpc.ServiceRegistrar, srv OAuthUserServiceServer) {
	// If the following call pancis, it indicates UnimplementedOAuthUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OAuthUserService_ServiceDesc, srv)
}

func _OAuthUserService_CreateOAuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOAuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthUserServiceServer).CreateOAuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuthUserService_CreateOAuthUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthUserServiceServer).CreateOAuthUser(ctx, req.(*CreateOAuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthUserService_DeleteOAuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOAuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthUserServiceServer).DeleteOAuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuthUserService_DeleteOAuthUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthUserServiceServer).DeleteOAuthUser(ctx, req.(*DeleteOAuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthUserService_SyncOAuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncOAuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthUserServiceServer).SyncOAuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuthUserService_SyncOAuthUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthUserServiceServer).SyncOAuthUser(ctx, req.(*SyncOAuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuthUserService_ServiceDesc is the grpc.ServiceDesc for OAuthUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuthUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.OAuthUserService",
	HandlerType: (*OAuthUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOAuthUser",
			Handler:    _OAuthUserService_CreateOAuthUser_Handler,
		},
		{
			MethodName: "DeleteOAuthUser",
			Handler:    _OAuthUserService_DeleteOAuthUser_Handler,
		},
		{
			MethodName: "SyncOAuthUser",
			Handler:    _OAuthUserService_SyncOAuthUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/user.proto",
}
