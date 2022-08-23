package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	//"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) AdmitPatient(ctx context.Context, req *pb.AdmitPatientParameters) (*pb.AddedAdmitPatientResult, error) {
	var err error
	ans := &pb.AddedAdmitPatientResult{}

	if req.PatientId == "" || req.CenterId == "" || req.CreatedAt == "" {
		if req.PatientId == "" {
			fmt.Println("Admit patient rpc error : patientid cannot be null")
		}
		if req.CenterId == "" {
			fmt.Println("Admit patient rpc error : centerid  cannot be null")
		}
		if req.CreatedAt == "" {
			fmt.Println("Admit patient rpc error : createdat cannot be null")
		}
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	// layout := "2006-01-02T15:04:05.000Z"
	// time, err := time.Parse(layout, req.ObservationDateTime)
	if err != nil {
		fmt.Println("AdmitPatient  rpc : ", err)
	}
	answer, err := grpca.admitpatientapi.AdmitPatientApi(req.PatientId, req.CenterId, req.CreatedAt)

	if err != nil {
		fmt.Println("AddObservation rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	ans = &pb.AddedAdmitPatientResult{
		Value: answer,
	}

	return ans, nil

}
