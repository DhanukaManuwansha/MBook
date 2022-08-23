package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lunux2008/xulu"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) AddObservation(ctx context.Context, req *pb.AddObservationParameters) (*pb.AddedObservationResult, error) {
	var err error
	ans := &pb.AddedObservationResult{}

	if req.ObservationDateTime == "" || req.PatientId == "" || req.NurseId == 0 {
		if req.ObservationDateTime == "" {
			fmt.Println("Add Observation rpc error : observation time cannot be null")
		}
		if req.PatientId == "" {
			fmt.Println("Add Observation rpc error : patientId cannot be null")
		}
		if req.NurseId == 0 {
			fmt.Println("Add observation rpc error : nurseId cannot be 0")
		}
		ans = &pb.AddedObservationResult{
			Message: "Server error",
			Status:  http.StatusBadRequest,
		}
		return ans, status.Error(codes.InvalidArgument, "missing required params")

	}
	if len(strings.TrimSpace(req.Notes)) >= 20 {
		ans = &pb.AddedObservationResult{
			Message: "Too many letters.",
			Status:  http.StatusBadRequest,
		}

		return ans, nil
	}
	fmt.Print(req.ObservationDateTime)

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.ObservationDateTime)
	if err != nil {
		ans = &pb.AddedObservationResult{
			Message: "Server error",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("AddObservation  rpc : ", err)
		return ans, err
	}
	answer, err := grpca.api.AddObservationApi(time, float64(req.BpSis), float64(req.BpDia), float64(req.Pr), float64(req.Rr), float64(req.Temp), req.Notes, req.PatientId, req.NurseId)

	if err != nil {
		ans = &pb.AddedObservationResult{
			Message: "Server error",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("AddObservation  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	ans = &pb.AddedObservationResult{
		Value:   answer,
		Message: "Observation successfully added.",
		Status:  http.StatusBadRequest,
	}

	return ans, nil

}

func (grpca GRPCAdapter) GetObservationById(ctx context.Context, req *pb.GetObservationById) (*pb.ObservationResult, error) {
	var err error
	ans := &pb.ObservationResult{}

	if req.ObseravationId == 0 {
		fmt.Println("Get observation by ID rpc error : observation ID cannot be 0")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	obsId, obsDateTime, obsBPSis, obsBPDia, obsPR, obsRR, obsTemp, obsNote, obsPatientId, obsNurseId, err := grpca.api.GetObservationByIDApi(req.ObseravationId)

	if err != nil {
		fmt.Println("GetObservationById  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	ans = &pb.ObservationResult{
		ObservationId:       obsId,
		ObservationDateTime: obsDateTime.String(),
		BpSis:               float32(obsBPSis),
		BpDia:               float32(obsBPDia),
		Pr:                  float32(obsPR),
		Rr:                  float32(obsRR),
		Temp:                float32(obsTemp),
		Notes:               obsNote,
		PatientId:           obsPatientId,
		NurseId:             obsNurseId,
	}

	return ans, nil

}

//Get all blood pressure values in rpc layer					-Dhanuka Manuwansha						-26/05/2022
func (grpca GRPCAdapter) GetAllBPSis(ctx context.Context, req *pb.PatientIdParameter) (*pb.GetAllBPResult, error) {
	var err error

	ans := &pb.GetAllBPResult{}
	//Check is patient Id null					-Dhanuka Manuwansha						-26/05/2022
	if req.PatientId == "" {
		fmt.Println("Get All BP rpc error : patient ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get data from api layer					-Dhanuka Manuwansha						-26/05/2022
	bpValues, err := grpca.api.GetAllBPSisValuesApi(req.PatientId)

	if err != nil {
		fmt.Println("GetAllBPValues  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	var tempBP []float32
	//Conert float64 to float32 					-Dhanuka Manuwansha						-26/05/2022
	for i := 0; i < len(bpValues); i++ {
		tempBP = append(tempBP, float32(bpValues[i]))
	}

	//Add data to pass it to the user					-Dhanuka Manuwansha						-26/05/2022
	ans = &pb.GetAllBPResult{
		BpValues: tempBP,
	}

	return ans, nil

}

//Get all blood pressure values in rpc layer					-Dhanuka Manuwansha						-26/05/2022
func (grpca GRPCAdapter) GetAllBPDia(ctx context.Context, req *pb.PatientIdParameter) (*pb.GetAllBPResult, error) {
	var err error

	ans := &pb.GetAllBPResult{}
	//Check is patient Id null					-Dhanuka Manuwansha						-26/05/2022
	if req.PatientId == "" {
		fmt.Println("Get All BP rpc error : patient ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get data from api layer					-Dhanuka Manuwansha						-26/05/2022
	bpValues, err := grpca.api.GetAllBPDiaValuesApi(req.PatientId)

	if err != nil {
		fmt.Println("GetAllBPValues  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	var tempBP []float32
	//Conert float64 to float32 					-Dhanuka Manuwansha						-26/05/2022
	for i := 0; i < len(bpValues); i++ {
		tempBP = append(tempBP, float32(bpValues[i]))
	}

	//Add data to pass it to the user					-Dhanuka Manuwansha						-26/05/2022
	ans = &pb.GetAllBPResult{
		BpValues: tempBP,
	}

	return ans, nil

}

//Get all pulse rate pressure values in rpc layer					-Dhanuka Manuwansha						-26/05/2022
func (grpca GRPCAdapter) GetAllPR(ctx context.Context, req *pb.PatientIdParameter) (*pb.GetAllPRResult, error) {
	var err error
	ans := &pb.GetAllPRResult{}

	if req.PatientId == "" {
		fmt.Println("Get All PR rpc error : Patient ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	prValues, err := grpca.api.GetAllPRValuesApi(req.PatientId)

	if err != nil {
		fmt.Println("GetAllPRValues  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	var tempPR []float32
	for i := 0; i < len(prValues); i++ {
		tempPR = append(tempPR, float32(prValues[i]))
	}

	ans = &pb.GetAllPRResult{
		PrValues: tempPR,
	}

	return ans, nil

}

//Get all respiratory rate values in rpc layer					-Dhanuka Manuwansha						-26/05/2022
func (grpca GRPCAdapter) GetAllRR(ctx context.Context, req *pb.PatientIdParameter) (*pb.GetAllRRResult, error) {
	var err error
	ans := &pb.GetAllRRResult{}

	if req.PatientId == "" {
		fmt.Println("Get All RR rpc error : Patient ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	rrValues, err := grpca.api.GetAllRRValuesApi(req.PatientId)

	if err != nil {
		fmt.Println("GetAllRRValue  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	var tempRR []float32
	for i := 0; i < len(rrValues); i++ {
		tempRR = append(tempRR, float32(rrValues[i]))
	}

	ans = &pb.GetAllRRResult{
		Rrvalues: tempRR,
	}

	return ans, nil

}

//Get all temperature values in rpc layer					-Dhanuka Manuwansha						-26/05/2022
func (grpca GRPCAdapter) GetAllTemp(ctx context.Context, req *pb.PatientIdParameter) (*pb.GetAllTempResult, error) {
	var err error
	ans := &pb.GetAllTempResult{}

	if req.PatientId == "" {

		fmt.Println("Get All Temp rpc error : Patient ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	tempValues, err := grpca.api.GetAllTempValuesApi(req.PatientId)

	if err != nil {
		fmt.Println("GetAllTempValues rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	var temp []float32
	for i := 0; i < len(tempValues); i++ {
		temp = append(temp, float32(tempValues[i]))
	}

	ans = &pb.GetAllTempResult{
		TempValues: temp,
	}

	return ans, nil

}

//Get all observations in rpc layer					-Dhanuka Manuwansha						-26/05/2022
func (grpca GRPCAdapter) GetAllObservations(ctx context.Context, req *pb.PatientIdParameter) (*pb.GetAllObservationsResult, error) {
	var err error
	ans := &pb.GetAllObservationsResult{}

	//Check patient Id is null					-Dhanuka Manuwansha						-26/05/2022
	if req.PatientId == "" {
		fmt.Println("Get All Observations rpc error : Patient ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get all observation from the api layer					-Dhanuka Manuwansha						-26/05/2022
	obs, err := grpca.api.GetAllObservationsApi(req.PatientId)

	if err != nil {
		fmt.Println("GetAllObservation  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")

	}
	//Tempory observation list to add all observations got from api layer					-Dhanuka Manuwansha						-26/05/2022
	var observationList []*pb.ObservationResult
	//Adding each observation to tempory object					-Dhanuka Manuwansha						-26/05/2022
	for i, v := range obs {
		var tempObservation = &pb.ObservationResult{
			ObservationId:       v.ObservationID,
			ObservationDateTime: v.ObsDatetime.String(),
			BpSis:               float32(v.BpSys),
			BpDia:               float32(v.BpDia),
			Pr:                  float32(v.Pr),
			Rr:                  float32(v.Rr),
			Temp:                float32(v.Temp),
			Notes:               v.Notes.String,
			PatientId:           v.PatientID,
			NurseId:             v.NurseID,
		}

		fmt.Println(i)
		//Add  each tempory object to final list					-Dhanuka Manuwansha						-26/05/2022
		observationList = append(observationList, tempObservation)
	}

	ans = &pb.GetAllObservationsResult{
		Observations: observationList,
	}

	return ans, nil

}

func (grpca GRPCAdapter) UpdateObservation(ctx context.Context, req *pb.UpdateObservationParameters) (*pb.UpdatedObservationResult, error) {
	var err error
	ans := &pb.UpdatedObservationResult{}

	//Check observation id is 0 or not					-Dhanuka Manuwansha						-27/05/2022
	if req.ObservationId == 0 {
		ans = &pb.UpdatedObservationResult{
			Message: "Server error",
			Status:  http.StatusBadRequest,
		}

		fmt.Println("Update observation rpc error : observation ID cannot be 0")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	if len(strings.TrimSpace(req.Notes)) >= 20 {
		ans = &pb.UpdatedObservationResult{
			Message: "Too many letters.",
			Status:  http.StatusBadRequest,
		}

		return ans, nil
	}

	fmt.Println(req.ObservationDateTime)

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.ObservationDateTime)
	fmt.Println(time)

	if err != nil {
		fmt.Println("ubdateObservation rpc : ", err)
		ans = &pb.UpdatedObservationResult{
			Message: "Server error",
			Status:  http.StatusBadRequest,
		}

		return ans, nil
	}

	obsId, obsDateTime, obsBPSis, obsBPDia, obsPR, obsRR, obsTemp, obsNote, obsPatientId, obsNurseId, err := grpca.api.UpdateObservationApi(req.ObservationId, time, float64(req.BpSis), float64(req.BpDia), float64(req.Pr), float64(req.Rr), float64(req.Temp), req.Notes, req.PatientId, req.NurseId)
	xulu.Use(obsId, obsDateTime, obsBPSis, obsBPDia, obsPR, obsRR, obsTemp, obsNote, obsPatientId, obsNurseId)
	if err != nil {
		ans = &pb.UpdatedObservationResult{
			Message: "Server error",
			Status:  http.StatusBadRequest,
		}

		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	ans = &pb.UpdatedObservationResult{
		Message: "Observation successfully updated.",
		Status:  http.StatusOK,
	}

	return ans, nil

}
