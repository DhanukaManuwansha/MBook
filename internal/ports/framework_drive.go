package ports

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
)

type GRPCPort interface {
	Run()
	// user
	RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetUserByNameResponse, error)
	GetUserByNIC(ctx context.Context, req *pb.GetUserByNICRequest) (*pb.GetUserByNICResponse, error)
	GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error)

	// ward
	GetWardsByCenter(ctx context.Context, req *pb.GetWardsByCenterRequest) (*pb.GetWardsByCenterResponse, error)

	// bed
	GetBedsByWard(ctx context.Context, req *pb.GetBedsbyWardRequest) (*pb.GetBedsbyWardResponse, error)

	// super admins
	RegisterSuperAdmin(ctx context.Context, req *pb.RegisterSuperAdminRequest) (*pb.RegisterSuperAdminResponse, error)
	// GetSuperAdmins(ctx context.Context, req *pb.GetSuperAdminsRequest) (*pb.GetSuperAdminsResponse, error)

	// center admins
	RegisterCenterAdmin(ctx context.Context, req *pb.RegisterCenterAdminRequest) (*pb.RegisterCenterAdminResponse, error)
	GetAllCenterAdmins(ctx context.Context, req *pb.GetAllCenterAdminsRequest) (*pb.GetAllCenterAdminsResponse, error)

	// nurse
	RegisterNurse(ctx context.Context, req *pb.RegisterNurseRequest) (*pb.RegisterNurseResponse, error)
	GetAllNurses(ctx context.Context, req *pb.GetAllNursesRequest) (*pb.GetAllNursesResponse, error)
	GetNurseFilter(ctx context.Context, req *pb.GetNurseFilterRequest) (*pb.GetAllNursesResponse, error)

	// doctor
	RegisterDoctor(ctx context.Context, req *pb.RegisterDoctorRequest) (*pb.RegisterDoctorResponse, error)
	AddObservation(ctx context.Context, req *pb.AddObservationParameters) (*pb.AddedObservationResult, error)
}
