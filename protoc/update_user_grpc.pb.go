// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: update_user.proto

package protoc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UpdateUserClient is the client API for UpdateUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateUserClient interface {
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	UpdateUserMember(ctx context.Context, in *UpdateUserMemberRequest, opts ...grpc.CallOption) (*UpdateUserMemberResponse, error)
	UpdateUserPassWord(ctx context.Context, in *UpdateUserPassWordRequest, opts ...grpc.CallOption) (*UpdateUserPassWordResponse, error)
}

type updateUserClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateUserClient(cc grpc.ClientConnInterface) UpdateUserClient {
	return &updateUserClient{cc}
}

func (c *updateUserClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/update_user.UpdateUser/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateUserClient) UpdateUserMember(ctx context.Context, in *UpdateUserMemberRequest, opts ...grpc.CallOption) (*UpdateUserMemberResponse, error) {
	out := new(UpdateUserMemberResponse)
	err := c.cc.Invoke(ctx, "/update_user.UpdateUser/UpdateUserMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateUserClient) UpdateUserPassWord(ctx context.Context, in *UpdateUserPassWordRequest, opts ...grpc.CallOption) (*UpdateUserPassWordResponse, error) {
	out := new(UpdateUserPassWordResponse)
	err := c.cc.Invoke(ctx, "/update_user.UpdateUser/UpdateUserPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateUserServer is the server API for UpdateUser service.
// All implementations must embed UnimplementedUpdateUserServer
// for forward compatibility
type UpdateUserServer interface {
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	UpdateUserMember(context.Context, *UpdateUserMemberRequest) (*UpdateUserMemberResponse, error)
	UpdateUserPassWord(context.Context, *UpdateUserPassWordRequest) (*UpdateUserPassWordResponse, error)
	mustEmbedUnimplementedUpdateUserServer()
}

// UnimplementedUpdateUserServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateUserServer struct {
}

func (UnimplementedUpdateUserServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUpdateUserServer) UpdateUserMember(context.Context, *UpdateUserMemberRequest) (*UpdateUserMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMember not implemented")
}
func (UnimplementedUpdateUserServer) UpdateUserPassWord(context.Context, *UpdateUserPassWordRequest) (*UpdateUserPassWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassWord not implemented")
}
func (UnimplementedUpdateUserServer) mustEmbedUnimplementedUpdateUserServer() {}

// UnsafeUpdateUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateUserServer will
// result in compilation errors.
type UnsafeUpdateUserServer interface {
	mustEmbedUnimplementedUpdateUserServer()
}

func RegisterUpdateUserServer(s grpc.ServiceRegistrar, srv UpdateUserServer) {
	s.RegisterService(&UpdateUser_ServiceDesc, srv)
}

func _UpdateUser_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateUserServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update_user.UpdateUser/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateUserServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateUser_UpdateUserMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateUserServer).UpdateUserMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update_user.UpdateUser/UpdateUserMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateUserServer).UpdateUserMember(ctx, req.(*UpdateUserMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateUser_UpdateUserPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPassWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateUserServer).UpdateUserPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update_user.UpdateUser/UpdateUserPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateUserServer).UpdateUserPassWord(ctx, req.(*UpdateUserPassWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateUser_ServiceDesc is the grpc.ServiceDesc for UpdateUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "update_user.UpdateUser",
	HandlerType: (*UpdateUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UpdateUser_UpdateUserInfo_Handler,
		},
		{
			MethodName: "UpdateUserMember",
			Handler:    _UpdateUser_UpdateUserMember_Handler,
		},
		{
			MethodName: "UpdateUserPassWord",
			Handler:    _UpdateUser_UpdateUserPassWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update_user.proto",
}