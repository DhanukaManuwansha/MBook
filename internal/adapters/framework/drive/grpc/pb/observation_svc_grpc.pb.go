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

// ObservationServiceClient is the client API for ObservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObservationServiceClient interface {
	AddObservation(ctx context.Context, in *AddObservationParameters, opts ...grpc.CallOption) (*AddedObservationResult, error)
	GetObservationById(ctx context.Context, in *GetObservationById, opts ...grpc.CallOption) (*ObservationResult, error)
	UpdateObservation(ctx context.Context, in *UpdateObservationParameters, opts ...grpc.CallOption) (*UpdatedObservationResult, error)
	GetAllObservations(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllObservationsResult, error)
	GetAllBPSis(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllBPResult, error)
	GetAllBPDia(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllBPResult, error)
	GetAllPR(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllPRResult, error)
	GetAllRR(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllRRResult, error)
	GetAllTemp(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllTempResult, error)
}

type observationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObservationServiceClient(cc grpc.ClientConnInterface) ObservationServiceClient {
	return &observationServiceClient{cc}
}

func (c *observationServiceClient) AddObservation(ctx context.Context, in *AddObservationParameters, opts ...grpc.CallOption) (*AddedObservationResult, error) {
	out := new(AddedObservationResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/AddObservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetObservationById(ctx context.Context, in *GetObservationById, opts ...grpc.CallOption) (*ObservationResult, error) {
	out := new(ObservationResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/GetObservationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) UpdateObservation(ctx context.Context, in *UpdateObservationParameters, opts ...grpc.CallOption) (*UpdatedObservationResult, error) {
	out := new(UpdatedObservationResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/UpdateObservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetAllObservations(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllObservationsResult, error) {
	out := new(GetAllObservationsResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/getAllObservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetAllBPSis(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllBPResult, error) {
	out := new(GetAllBPResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/getAllBPSis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetAllBPDia(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllBPResult, error) {
	out := new(GetAllBPResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/getAllBPDia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetAllPR(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllPRResult, error) {
	out := new(GetAllPRResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/getAllPR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetAllRR(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllRRResult, error) {
	out := new(GetAllRRResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/getAllRR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observationServiceClient) GetAllTemp(ctx context.Context, in *PatientIdParameter, opts ...grpc.CallOption) (*GetAllTempResult, error) {
	out := new(GetAllTempResult)
	err := c.cc.Invoke(ctx, "/pb.ObservationService/getAllTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObservationServiceServer is the server API for ObservationService service.
// All implementations should embed UnimplementedObservationServiceServer
// for forward compatibility
type ObservationServiceServer interface {
	AddObservation(context.Context, *AddObservationParameters) (*AddedObservationResult, error)
	GetObservationById(context.Context, *GetObservationById) (*ObservationResult, error)
	UpdateObservation(context.Context, *UpdateObservationParameters) (*UpdatedObservationResult, error)
	GetAllObservations(context.Context, *PatientIdParameter) (*GetAllObservationsResult, error)
	GetAllBPSis(context.Context, *PatientIdParameter) (*GetAllBPResult, error)
	GetAllBPDia(context.Context, *PatientIdParameter) (*GetAllBPResult, error)
	GetAllPR(context.Context, *PatientIdParameter) (*GetAllPRResult, error)
	GetAllRR(context.Context, *PatientIdParameter) (*GetAllRRResult, error)
	GetAllTemp(context.Context, *PatientIdParameter) (*GetAllTempResult, error)
}

// UnimplementedObservationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedObservationServiceServer struct {
}

func (UnimplementedObservationServiceServer) AddObservation(context.Context, *AddObservationParameters) (*AddedObservationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddObservation not implemented")
}
func (UnimplementedObservationServiceServer) GetObservationById(context.Context, *GetObservationById) (*ObservationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObservationById not implemented")
}
func (UnimplementedObservationServiceServer) UpdateObservation(context.Context, *UpdateObservationParameters) (*UpdatedObservationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObservation not implemented")
}
func (UnimplementedObservationServiceServer) GetAllObservations(context.Context, *PatientIdParameter) (*GetAllObservationsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllObservations not implemented")
}
func (UnimplementedObservationServiceServer) GetAllBPSis(context.Context, *PatientIdParameter) (*GetAllBPResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBPSis not implemented")
}
func (UnimplementedObservationServiceServer) GetAllBPDia(context.Context, *PatientIdParameter) (*GetAllBPResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBPDia not implemented")
}
func (UnimplementedObservationServiceServer) GetAllPR(context.Context, *PatientIdParameter) (*GetAllPRResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPR not implemented")
}
func (UnimplementedObservationServiceServer) GetAllRR(context.Context, *PatientIdParameter) (*GetAllRRResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRR not implemented")
}
func (UnimplementedObservationServiceServer) GetAllTemp(context.Context, *PatientIdParameter) (*GetAllTempResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTemp not implemented")
}

// UnsafeObservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObservationServiceServer will
// result in compilation errors.
type UnsafeObservationServiceServer interface {
	mustEmbedUnimplementedObservationServiceServer()
}

func RegisterObservationServiceServer(s grpc.ServiceRegistrar, srv ObservationServiceServer) {
	s.RegisterService(&ObservationService_ServiceDesc, srv)
}

func _ObservationService_AddObservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddObservationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).AddObservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/AddObservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).AddObservation(ctx, req.(*AddObservationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetObservationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObservationById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetObservationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/GetObservationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetObservationById(ctx, req.(*GetObservationById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_UpdateObservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObservationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).UpdateObservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/UpdateObservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).UpdateObservation(ctx, req.(*UpdateObservationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetAllObservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientIdParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetAllObservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/getAllObservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetAllObservations(ctx, req.(*PatientIdParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetAllBPSis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientIdParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetAllBPSis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/getAllBPSis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetAllBPSis(ctx, req.(*PatientIdParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetAllBPDia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientIdParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetAllBPDia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/getAllBPDia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetAllBPDia(ctx, req.(*PatientIdParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetAllPR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientIdParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetAllPR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/getAllPR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetAllPR(ctx, req.(*PatientIdParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetAllRR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientIdParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetAllRR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/getAllRR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetAllRR(ctx, req.(*PatientIdParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObservationService_GetAllTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientIdParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservationServiceServer).GetAllTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ObservationService/getAllTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservationServiceServer).GetAllTemp(ctx, req.(*PatientIdParameter))
	}
	return interceptor(ctx, in, info, handler)
}

// ObservationService_ServiceDesc is the grpc.ServiceDesc for ObservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ObservationService",
	HandlerType: (*ObservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddObservation",
			Handler:    _ObservationService_AddObservation_Handler,
		},
		{
			MethodName: "GetObservationById",
			Handler:    _ObservationService_GetObservationById_Handler,
		},
		{
			MethodName: "UpdateObservation",
			Handler:    _ObservationService_UpdateObservation_Handler,
		},
		{
			MethodName: "getAllObservations",
			Handler:    _ObservationService_GetAllObservations_Handler,
		},
		{
			MethodName: "getAllBPSis",
			Handler:    _ObservationService_GetAllBPSis_Handler,
		},
		{
			MethodName: "getAllBPDia",
			Handler:    _ObservationService_GetAllBPDia_Handler,
		},
		{
			MethodName: "getAllPR",
			Handler:    _ObservationService_GetAllPR_Handler,
		},
		{
			MethodName: "getAllRR",
			Handler:    _ObservationService_GetAllRR_Handler,
		},
		{
			MethodName: "getAllTemp",
			Handler:    _ObservationService_GetAllTemp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "observation_svc.proto",
}
