package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetWardsByCenter(ctx context.Context, req *pb.GetWardsByCenterRequest) (*pb.GetWardsByCenterResponse, error) {

	var err error
	ans := &pb.GetWardsByCenterResponse{}

	if req.CenterId == 0 {
		return ans, status.Error(codes.InvalidArgument, "Missing required arguments")
	}

	wards, err := grpca.wardapi.GetWardsByCenterAPI(req.CenterId)

	if err != nil {
		fmt.Printf("Something wrong with the wardRPC: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error!")
	}

	var wardList []*pb.Ward

	for i := 0; i < len(wards); i++ {
		var tempWards = &pb.Ward{
			WardNo:       wards[i].WardNo,
			WardId:       wards[i].WardID,
			TotOfPatient: wards[i].TotOfPatient,
			WardType:     wards[i].WardType,
			Period:       wards[i].Period,
			CreatedAt:    wards[i].CreatedAt.String(),
			CenterId:     wards[i].CenterID,
		}

		wardList = append(wardList, tempWards)
	}

	ans = &pb.GetWardsByCenterResponse{
		Wards: wardList,
	}

	return ans, nil
}
