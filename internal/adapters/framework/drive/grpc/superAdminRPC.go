package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) RegisterSuperAdmin(ctx context.Context, req *pb.RegisterSuperAdminRequest) (*pb.RegisterSuperAdminResponse, error) {
	var err error
	ans := &pb.RegisterSuperAdminResponse{}

	newId := uuid.New()
	id := newId.String()

	superAdmin, err := grpca.userapi.RegisterUserAPI(id, req.UserName, req.FirstName, req.LastName, req.Nic, req.TellNo, req.Address, req.UserEmail, req.UserPwd, req.IsEmailVerified, req.IsTellNoVerified)
	superAdminId, userId, err := grpca.superadminapi.RegisterSuperAdminAPI(superAdmin)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error from api")
	}

	ans = &pb.RegisterSuperAdminResponse{
		SuperAdminId: superAdminId,
		UserId:       userId,
	}

	return ans, nil

}

// func (grpca GRPCAdapter) GetSuperAdmins(ctx context.Context, req *pb.GetSuperAdminsRequest) (*pb.GetSuperAdminsResponse, error) {
// 	var err error
// 	ans := &pb.GetSuperAdminsResponse{}

// 	fmt.Println("2")

// 	superAdmins, err := grpca.superadminapi.GetSuperAdmins()

// 	fmt.Println("3")

// 	if err != nil {
// 		fmt.Printf("error: %v", err)
// 		return ans, status.Error(codes.Internal, "Unexpected Error")
// 	}

// 	var superAdminList []*pb.SuperAdmin

// 	for i := 0; i < len(superAdmins); i++ {
// 		var tempSuperAdmins = &pb.User{
// 			UserId:    superAdmins[i].UserID,

// 		}

// 		superAdminList = append(superAdminList, tempUsers)
// 	}

// 	ans = &pb.GetSuperAdminsResponse{
// 		SuperAdmins: superAdminList,
// 	}

// 	return ans, nil
// }
