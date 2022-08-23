package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetAllCenters(context.Context, *pb.GetAllCetersRequest) (*pb.GetAllCetersResponse, error) {
	var err error
	results := &pb.GetAllCetersResponse{}

	centers, err := grpca.centerapi.GetAllCenterIds()

	if err != nil {
		fmt.Println("GetAllNurseNotes rpc error : ", err)
		return results, status.Error(codes.Internal, "unexpected Error")
	}

	var centersList []*pb.Center
	for i, v := range centers {
		var tempCenter = &pb.Center{
			CenterId:     v.CenterID,
			CenterName:   v.CenterName,
			Address:      v.Address.String,
			TellNo:       v.TellNo,
			Email:        v.Email,
			WebSite:      v.Website,
			ActiveStatus: v.ActiveStatus,
			OnlineStatus: v.OnlineStatus,
			CenterTypeId: v.CentertypeID,
		}

		fmt.Println(i)

		centersList = append(centersList, tempCenter)
	}

	results = &pb.GetAllCetersResponse{
		Centers: centersList,
	}

	return results, nil

}
