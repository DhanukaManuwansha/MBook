package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetBedsByWard(ctx context.Context, req *pb.GetBedsbyWardRequest) (*pb.GetBedsbyWardResponse, error) {

	var err error
	ans := &pb.GetBedsbyWardResponse{}

	if req.WardId == 0 {
		return ans, status.Error(codes.InvalidArgument, "Missing required arguments")
	}

	beds, err := grpca.bedapi.GetBedsByWardApi(req.WardId)

	if err != nil {
		fmt.Printf("Something wrong with the bedRPC: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error!")
	}

	var tempBeds []*pb.GetBedsbyWardResult
	for i := 0; i < len(beds); i++ {
		var tempBed = &pb.GetBedsbyWardResult{
			BedTicketId: beds[i].BedTicketID,
			BedNo:       beds[i].BedNo,
			PatientId:   beds[i].PatientID,
		}

		tempBeds = append(tempBeds, tempBed)
	}

	ans = &pb.GetBedsbyWardResponse{
		BedTicket: tempBeds,
	}

	return ans, nil
}
