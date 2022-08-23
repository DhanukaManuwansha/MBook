package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetLatestPrescriptions(ctx context.Context, req *pb.GetLatestPrescriptionsParameters) (*pb.LatestPrescriptionResult, error) {
	var err error
	pres := &pb.LatestPrescriptionResult{}

	//Check is patient Id null or not					-Dhanuka Manuwansha						-27/05/2022
	if req.PatientId == "" {
		fmt.Println("Get latest prescription rpc error : patient ID cannot be null")
		return pres, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get latest prescription from API layer					-Dhanuka Manuwansha						-27/05/2022
	prescriptions, err := grpca.prescriptionapi.GetActivePrescriptionsApi(req.PatientId)

	if err != nil {
		fmt.Println("GetLatestePrescriptions rpc error : ", err)
		return pres, status.Error(codes.Internal, "unexpected Error")
	}

	var p []*pb.PrescriptionResult

	for i, v := range prescriptions {
		var tempPres = &pb.PrescriptionResult{
			PrescriptionId: v.PrescriptionID,
			PresDate:       v.PresDate.String(),
			PresTime:       v.PresTime.String(),
			ActiveStatus:   v.ActiveStatus,
			Notes:          v.Notes.String,
			Sketch:         v.Sketch.String,
			PatientId:      v.PatientID,
			DoctorId:       v.DoctorID,
		}

		fmt.Println(i)

		p = append(p, tempPres)
	}

	pres = &pb.LatestPrescriptionResult{
		Prescriptions: p,
	}

	return pres, nil

}
