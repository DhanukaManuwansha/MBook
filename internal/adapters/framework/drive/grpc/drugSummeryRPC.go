package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetDrugSummery(ctx context.Context, req *pb.GetDrugSummeryParameters) (*pb.DrugSummeryListResult, error) {
	var err error
	drugsum := &pb.DrugSummeryListResult{}

	//Check is patient Id null or not					-Dhanuka Manuwansha						-27/05/2022
	if req.PatientId == "" {
		fmt.Println("Get Drug summery rpc error : patient ID cannot be null")
		return drugsum, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get drug summery from API layer					-Dhanuka Manuwansha						-27/05/2022
	drugsummery, err := grpca.drugsummeryapi.GetAllDrugSummeryApi(req.PatientId)

	if err != nil {
		fmt.Println("GetDrugChart rpc error : ", err)
		return drugsum, status.Error(codes.Internal, "unexpected Error")
	}

	var drugList []*pb.DrugSummeryResult

	for i, v := range drugsummery {
		var tempDrug = &pb.DrugSummeryResult{
			DrugsummeryId: v.DrugsummeryID,
			SummeryDate:   v.SummeryDate.String(),
			SummeryStatus: v.SummeryStatus,
			DrugorderId:   int64(v.DrugorderID.Int64),
			PatientId:     v.PatientID,
		}

		fmt.Println(i)

		drugList = append(drugList, tempDrug)
	}

	drugsum = &pb.DrugSummeryListResult{
		Drugsummery: drugList,
	}

	return drugsum, nil

}

func (grpca GRPCAdapter) GetDrugSummeryOfADrug(ctx context.Context, req *pb.DrugSummeryOfADrugResultParam) (*pb.DrugSummeryOfADrugResultList, error) {
	var err error
	drugsum := &pb.DrugSummeryOfADrugResultList{}

	//Check is patient Id null or not					-Dhanuka Manuwansha						-27/05/2022
	if req.PatientId == "" {
		fmt.Println("Get Drug summery rpc error : patient ID cannot be null")
		return drugsum, status.Error(codes.InvalidArgument, "missing required params")
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	endDate, err := time.Parse("2006-01-02", req.EndDate)

	var param = db.GetAllDrugSummeryOfADrugParams{
		PatientID:     req.PatientId,
		SummeryDate:   startDate,
		SummeryDate_2: endDate,
	}

	//Get drug summery from API layer					-Dhanuka Manuwansha						-27/05/2022
	drugsummery, err := grpca.drugsummeryapi.GetAllDrugSummeryOfADrugApi(param)

	if err != nil {
		fmt.Println("GetDrugChart rpc error : ", err)
		return drugsum, status.Error(codes.Internal, "unexpected Error")
	}

	var drugList []*pb.DrugSummeryOfADrugResult

	for i, v := range drugsummery {
		var tempDrug = &pb.DrugSummeryOfADrugResult{
			DrugName:         v.DrugName,
			Dose:             v.Dose,
			Dosage:           v.Dosage,
			Frequency:        v.Frequency,
			GivenDate:        v.Givendate.String(),
			GiveUntill:       v.Giveuntil.String(),
			DrugsummeryId:    v.DrugsummeryID,
			SummeryDate:      v.SummeryDate.String(),
			SummeryStatus:    v.SummeryStatus,
			FirstDoseStatus:  v.FirstDoseIsGiven,
			SecondDoseStatus: v.SecondDoseIsGiven,
			ThirdDoseStatus:  v.ThirdDoseIsGiven,
			FourthDoseStatus: v.FourthDoseIsGiven,
			DrugorderId:      int64(v.DrugorderID.Int64),
			PatientId:        v.PatientID,
		}

		fmt.Println(i)

		drugList = append(drugList, tempDrug)
	}

	drugsum = &pb.DrugSummeryOfADrugResultList{
		Drugsummery: drugList,
	}

	return drugsum, nil

}

func (grpca GRPCAdapter) UpdateDrugSummery(ctx context.Context, req *pb.DrugSummeryUpdateParam) (*pb.DrugSummeryUpdateResult, error) {
	var err error
	ans := &pb.DrugSummeryUpdateResult{}

	if req.DrugsummeryId == 0 {
		fmt.Println("Update Drug Summery rpc error : DrugSummery ID cannot be 0")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.SummeryDate)
	fmt.Println(time)

	if err != nil {
		fmt.Println("ubdateDrugSummery rpc : ", err)
	}

	var param = db.UpdateDrugSummeryParams{
		DrugsummeryID:     req.DrugsummeryId,
		SummeryDate:       time,
		SummeryStatus:     req.SummeryStatus,
		FirstDoseIsGiven:  req.FirstDoseStatus,
		SecondDoseIsGiven: req.SecondDoseStatus,
		ThirdDoseIsGiven:  req.ThirdDoseStatus,
		FourthDoseIsGiven: req.FourthDoseStatus,
		DrugorderID:       sql.NullInt64{Valid: true, Int64: req.DrugorderId},
		PatientID:         req.PatientId,
	}

	updatedSummery, err := grpca.drugsummeryapi.UpdateDrugSummeryApi(param)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	ans = &pb.DrugSummeryUpdateResult{
		DrugsummeryId:    updatedSummery.DrugsummeryID,
		SummeryDate:      updatedSummery.SummeryDate.String(),
		SummeryStatus:    updatedSummery.SummeryStatus,
		FirstDoseStatus:  updatedSummery.FirstDoseIsGiven,
		SecondDoseStatus: updatedSummery.SecondDoseIsGiven,
		ThirdDoseStatus:  updatedSummery.ThirdDoseIsGiven,
		FourthDoseStatus: updatedSummery.FourthDoseIsGiven,
		DrugorderId:      int64(updatedSummery.DrugorderID.Int64),
		PatientId:        updatedSummery.PatientID,
	}

	return ans, nil

}
