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

// DrugChartServiceClient is the client API for DrugChartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DrugChartServiceClient interface {
	GetDrugChart(ctx context.Context, in *GetDrugChartParameters, opts ...grpc.CallOption) (*DrugChartResult, error)
	GetDrugChartList(ctx context.Context, in *GetDrugChartParameters, opts ...grpc.CallOption) (*DrugChartResultForDashBoard, error)
	GetDrugChartListForNurseDesk(ctx context.Context, in *DrugChartForNurseDesktParam, opts ...grpc.CallOption) (*DrugChartForNurseDeskResultList, error)
}

type drugChartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDrugChartServiceClient(cc grpc.ClientConnInterface) DrugChartServiceClient {
	return &drugChartServiceClient{cc}
}

func (c *drugChartServiceClient) GetDrugChart(ctx context.Context, in *GetDrugChartParameters, opts ...grpc.CallOption) (*DrugChartResult, error) {
	out := new(DrugChartResult)
	err := c.cc.Invoke(ctx, "/pb.DrugChartService/GetDrugChart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drugChartServiceClient) GetDrugChartList(ctx context.Context, in *GetDrugChartParameters, opts ...grpc.CallOption) (*DrugChartResultForDashBoard, error) {
	out := new(DrugChartResultForDashBoard)
	err := c.cc.Invoke(ctx, "/pb.DrugChartService/GetDrugChartList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drugChartServiceClient) GetDrugChartListForNurseDesk(ctx context.Context, in *DrugChartForNurseDesktParam, opts ...grpc.CallOption) (*DrugChartForNurseDeskResultList, error) {
	out := new(DrugChartForNurseDeskResultList)
	err := c.cc.Invoke(ctx, "/pb.DrugChartService/GetDrugChartListForNurseDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DrugChartServiceServer is the server API for DrugChartService service.
// All implementations should embed UnimplementedDrugChartServiceServer
// for forward compatibility
type DrugChartServiceServer interface {
	GetDrugChart(context.Context, *GetDrugChartParameters) (*DrugChartResult, error)
	GetDrugChartList(context.Context, *GetDrugChartParameters) (*DrugChartResultForDashBoard, error)
	GetDrugChartListForNurseDesk(context.Context, *DrugChartForNurseDesktParam) (*DrugChartForNurseDeskResultList, error)
}

// UnimplementedDrugChartServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDrugChartServiceServer struct {
}

func (UnimplementedDrugChartServiceServer) GetDrugChart(context.Context, *GetDrugChartParameters) (*DrugChartResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrugChart not implemented")
}
func (UnimplementedDrugChartServiceServer) GetDrugChartList(context.Context, *GetDrugChartParameters) (*DrugChartResultForDashBoard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrugChartList not implemented")
}
func (UnimplementedDrugChartServiceServer) GetDrugChartListForNurseDesk(context.Context, *DrugChartForNurseDesktParam) (*DrugChartForNurseDeskResultList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrugChartListForNurseDesk not implemented")
}

// UnsafeDrugChartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DrugChartServiceServer will
// result in compilation errors.
type UnsafeDrugChartServiceServer interface {
	mustEmbedUnimplementedDrugChartServiceServer()
}

func RegisterDrugChartServiceServer(s grpc.ServiceRegistrar, srv DrugChartServiceServer) {
	s.RegisterService(&DrugChartService_ServiceDesc, srv)
}

func _DrugChartService_GetDrugChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDrugChartParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DrugChartServiceServer).GetDrugChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DrugChartService/GetDrugChart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DrugChartServiceServer).GetDrugChart(ctx, req.(*GetDrugChartParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _DrugChartService_GetDrugChartList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDrugChartParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DrugChartServiceServer).GetDrugChartList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DrugChartService/GetDrugChartList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DrugChartServiceServer).GetDrugChartList(ctx, req.(*GetDrugChartParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _DrugChartService_GetDrugChartListForNurseDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrugChartForNurseDesktParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DrugChartServiceServer).GetDrugChartListForNurseDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DrugChartService/GetDrugChartListForNurseDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DrugChartServiceServer).GetDrugChartListForNurseDesk(ctx, req.(*DrugChartForNurseDesktParam))
	}
	return interceptor(ctx, in, info, handler)
}

// DrugChartService_ServiceDesc is the grpc.ServiceDesc for DrugChartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DrugChartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DrugChartService",
	HandlerType: (*DrugChartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDrugChart",
			Handler:    _DrugChartService_GetDrugChart_Handler,
		},
		{
			MethodName: "GetDrugChartList",
			Handler:    _DrugChartService_GetDrugChartList_Handler,
		},
		{
			MethodName: "GetDrugChartListForNurseDesk",
			Handler:    _DrugChartService_GetDrugChartListForNurseDesk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drugchart_svc.proto",
}