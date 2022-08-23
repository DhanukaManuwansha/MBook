// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// NurseServicesClient is the client API for NurseServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NurseServicesClient interface {
	RegisterNurse(ctx context.Context, in *RegisterNurseRequest, opts ...grpc.CallOption) (*RegisterNurseResponse, error)
	GetAllNurses(ctx context.Context, in *GetAllNursesRequest, opts ...grpc.CallOption) (*GetAllNursesResponse, error)
	GetNurseFilter(ctx context.Context, in *GetNurseFilterRequest, opts ...grpc.CallOption) (*GetAllNursesResponse, error)
	UpdateNurse(ctx context.Context, in *UpdateNurseRequest, opts ...grpc.CallOption) (*UpdateNurseResponse, error)
}

type nurseServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewNurseServicesClient(cc grpc.ClientConnInterface) NurseServicesClient {
	return &nurseServicesClient{cc}
}

func (c *nurseServicesClient) RegisterNurse(ctx context.Context, in *RegisterNurseRequest, opts ...grpc.CallOption) (*RegisterNurseResponse, error) {
	out := new(RegisterNurseResponse)
	err := c.cc.Invoke(ctx, "/pb.NurseServices/RegisterNurse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nurseServicesClient) GetAllNurses(ctx context.Context, in *GetAllNursesRequest, opts ...grpc.CallOption) (*GetAllNursesResponse, error) {
	out := new(GetAllNursesResponse)
	err := c.cc.Invoke(ctx, "/pb.NurseServices/GetAllNurses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nurseServicesClient) GetNurseFilter(ctx context.Context, in *GetNurseFilterRequest, opts ...grpc.CallOption) (*GetAllNursesResponse, error) {
	out := new(GetAllNursesResponse)
	err := c.cc.Invoke(ctx, "/pb.NurseServices/GetNurseFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nurseServicesClient) UpdateNurse(ctx context.Context, in *UpdateNurseRequest, opts ...grpc.CallOption) (*UpdateNurseResponse, error) {
	out := new(UpdateNurseResponse)
	err := c.cc.Invoke(ctx, "/pb.NurseServices/UpdateNurse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NurseServicesServer is the server API for NurseServices service.
// All implementations should embed UnimplementedNurseServicesServer
// for forward compatibility
type NurseServicesServer interface {
	RegisterNurse(context.Context, *RegisterNurseRequest) (*RegisterNurseResponse, error)
	GetAllNurses(context.Context, *GetAllNursesRequest) (*GetAllNursesResponse, error)
	GetNurseFilter(context.Context, *GetNurseFilterRequest) (*GetAllNursesResponse, error)
	UpdateNurse(context.Context, *UpdateNurseRequest) (*UpdateNurseResponse, error)
}

// UnimplementedNurseServicesServer should be embedded to have forward compatible implementations.
type UnimplementedNurseServicesServer struct {
}

func (UnimplementedNurseServicesServer) RegisterNurse(context.Context, *RegisterNurseRequest) (*RegisterNurseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNurse not implemented")
}
func (UnimplementedNurseServicesServer) GetAllNurses(context.Context, *GetAllNursesRequest) (*GetAllNursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNurses not implemented")
}
func (UnimplementedNurseServicesServer) GetNurseFilter(context.Context, *GetNurseFilterRequest) (*GetAllNursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNurseFilter not implemented")
}
func (UnimplementedNurseServicesServer) UpdateNurse(context.Context, *UpdateNurseRequest) (*UpdateNurseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNurse not implemented")
}

// UnsafeNurseServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NurseServicesServer will
// result in compilation errors.
type UnsafeNurseServicesServer interface {
	mustEmbedUnimplementedNurseServicesServer()
}

func RegisterNurseServicesServer(s grpc.ServiceRegistrar, srv NurseServicesServer) {
	s.RegisterService(&NurseServices_ServiceDesc, srv)
}

func _NurseServices_RegisterNurse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNurseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NurseServicesServer).RegisterNurse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NurseServices/RegisterNurse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NurseServicesServer).RegisterNurse(ctx, req.(*RegisterNurseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NurseServices_GetAllNurses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NurseServicesServer).GetAllNurses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NurseServices/GetAllNurses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NurseServicesServer).GetAllNurses(ctx, req.(*GetAllNursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NurseServices_GetNurseFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNurseFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NurseServicesServer).GetNurseFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NurseServices/GetNurseFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NurseServicesServer).GetNurseFilter(ctx, req.(*GetNurseFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NurseServices_UpdateNurse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNurseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NurseServicesServer).UpdateNurse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NurseServices/UpdateNurse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NurseServicesServer).UpdateNurse(ctx, req.(*UpdateNurseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NurseServices_ServiceDesc is the grpc.ServiceDesc for NurseServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NurseServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NurseServices",
	HandlerType: (*NurseServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNurse",
			Handler:    _NurseServices_RegisterNurse_Handler,
		},
		{
			MethodName: "GetAllNurses",
			Handler:    _NurseServices_GetAllNurses_Handler,
		},
		{
			MethodName: "GetNurseFilter",
			Handler:    _NurseServices_GetNurseFilter_Handler,
		},
		{
			MethodName: "UpdateNurse",
			Handler:    _NurseServices_UpdateNurse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nurse_svc.proto",
}