package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lunux2008/xulu"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Update Nurse Note in rpc layer					-Dhanuka Manuwansha						-27/05/2022
func (grpca GRPCAdapter) UpdateNurseNote(ctx context.Context, req *pb.UpdateNurseNoteParameters) (*pb.UpdatedNurseNoteResult, error) {
	var err error
	nurseNote := &pb.UpdatedNurseNoteResult{}

	//Check nurse note id is 0 or not					-Dhanuka Manuwansha						-27/05/2022
	if req.NursingNoteId == 0 {
		nurseNote = &pb.UpdatedNurseNoteResult{
			Message: "Server error, Try again later.",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("Update NurseNote rpc error : nursingnote ID cannot be 0")
		return nurseNote, status.Error(codes.InvalidArgument, "missing required params")
	}

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.NotesDateTimer)
	if err != nil {
		nurseNote = &pb.UpdatedNurseNoteResult{
			Message: "Server error, Try again later.",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("UpdateNurseNote rpc error : ", err)
		fmt.Println(err)
		return nurseNote, err
	}

	//Get data from API layer					-Dhanuka Manuwansha						-27/05/2022
	nursingNoteId, remark, nursingCare, notesDateTimer, nurse_id, patient_id, err := grpca.nursenoteapi.UpdateNurseNoteApi(req.NursingNoteId, req.Remark, req.NursingCare, time, req.NurseId, req.PatientId)
	xulu.Use(nursingNoteId, remark, nursingCare, notesDateTimer, nurse_id, patient_id)
	if err != nil {
		nurseNote = &pb.UpdatedNurseNoteResult{
			Message: "Server error, Try again later.",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("UpdateNurseNote rpc error : ", err)
		return nurseNote, status.Error(codes.Internal, "unexpected Error")
	}
	//Adding values to return it to the user					-Dhanuka Manuwansha						-27/05/2022
	nurseNote = &pb.UpdatedNurseNoteResult{
		Message: "Note successfully updated.",
		Status:  http.StatusOK,
	}

	return nurseNote, nil

}

//Get all nursenotes 					-Dhanuka Manuwansha						-27/05/2022
func (grpca GRPCAdapter) GetAllNurseNotes(ctx context.Context, req *pb.PatientIdParam) (*pb.GetAllNurseNotesResult, error) {
	var err error
	nurseNotes := &pb.GetAllNurseNotesResult{}

	//Check is patient Id null or not					-Dhanuka Manuwansha						-27/05/2022
	if req.PatientId == "" {
		fmt.Println("Get All NurseNotes rpc error : patient ID cannot be null")
		return nurseNotes, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get all nursenotes from API layer					-Dhanuka Manuwansha						-27/05/2022
	notes, err := grpca.nursenoteapi.GetAllNurseNotesApi(req.PatientId)
	fmt.Println(notes)
	if err != nil {
		fmt.Println("GetAllNurseNotes rpc error : ", err)
		return nurseNotes, status.Error(codes.Internal, "unexpected Error")
	}

	var nurseNoteList []*pb.NurseNoteResult
	for i, v := range notes {
		var tempNote = &pb.NurseNoteResult{
			NursingNoteId:  v.NursingnotesID,
			Remark:         v.Remark,
			NursingCare:    v.NursingCare.String,
			NotesDateTimer: v.NotesDatetime.Time.String(),
			PatientId:      v.PatientID,
			NurseId:        v.NurseID,
		}

		fmt.Println(i)
		fmt.Println(tempNote)

		nurseNoteList = append(nurseNoteList, tempNote)
	}

	nurseNotes = &pb.GetAllNurseNotesResult{
		NurseNotes: nurseNoteList,
	}

	return nurseNotes, nil

}

func (grpca GRPCAdapter) AddNurseNote(ctx context.Context, req *pb.AddNurseNoteParameters) (*pb.AddedNurseNoteResults, error) {
	var err error
	ans := &pb.AddedNurseNoteResults{}

	if req.Remark == "" || req.NotesDateTimer == "" || req.NurseId == "" {
		ans = &pb.AddedNurseNoteResults{
			Message: "Server error, Try again later.",
			Status:  http.StatusBadRequest,
		}
		if req.NotesDateTimer == "" {
			fmt.Println("Add NurseNote rpc error : Notes date cannot be null")
		}
		if req.Remark == "" {
			fmt.Println("Add NurseNote rpc error : Remark cannot be null")
		}
		if req.NurseId == "" {
			fmt.Println("Add NurseNote rpc error : NurseId cannot be null")
		}
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	// layout := "2006-01-02T15:04:05.000Z"
	time := time.Now()

	if err != nil {
		ans = &pb.AddedNurseNoteResults{
			Message: "Server error, Try again later.",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("AddNurseNote  rpc : ", err)
		return ans, err
	}
	answer, err := grpca.nursenoteapi.AddNurseNoteApi(req.Remark, req.NursingCare, req.NotesDateTimer, time.GoString(), req.NurseId, req.PatientId)

	if err != nil {
		ans = &pb.AddedNurseNoteResults{
			Message: "Server error, Try again later.",
			Status:  http.StatusBadRequest,
		}
		fmt.Println("AddNotes  rpc : ", err)
		return ans, status.Error(codes.Internal, "unexpected Error")
	}

	ans = &pb.AddedNurseNoteResults{
		Value:   answer,
		Message: "Note successfully updated",
		Status:  http.StatusOK,
	}

	return ans, nil
}

//Get all latest nursenotes 					-Dhanuka Manuwansha						-07/06/2022
func (grpca GRPCAdapter) GetLatestNurseNotes(ctx context.Context, req *pb.PatientIdParam) (*pb.GetAllLatestNurseNotesResult, error) {
	var err error
	nurseNotes := &pb.GetAllLatestNurseNotesResult{}

	//Check is patient Id null or not					-Dhanuka Manuwansha						-07/06/2022
	if req.PatientId == "" {
		fmt.Println("Get All Latest NurseNotes rpc error : patient ID cannot be null")
		return nurseNotes, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get all Latest nursenotes from API layer					-Dhanuka Manuwansha						-07/06/2022
	notes, err := grpca.nursenoteapi.GetLatestNurseNotesApi(req.PatientId)
	fmt.Println(notes)
	if err != nil {
		fmt.Println("GetAllLatestNurseNotes rpc error : ", err)
		return nurseNotes, status.Error(codes.Internal, "unexpected Error")
	}

	var nurseNoteList []*pb.LatestNurseNoteResult
	for i, v := range notes {
		var tempNote = &pb.LatestNurseNoteResult{

			Remark:         v.Remark,
			NursingCare:    v.NursingCare.String,
			NotesDateTimer: v.NotesDatetime.Time.String(),
			NurseName:      v.UserName,
		}

		fmt.Println(i)

		nurseNoteList = append(nurseNoteList, tempNote)
	}

	nurseNotes = &pb.GetAllLatestNurseNotesResult{
		NurseNotes: nurseNoteList,
	}

	return nurseNotes, nil

}
