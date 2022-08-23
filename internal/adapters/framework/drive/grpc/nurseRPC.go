package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) RegisterNurse(ctx context.Context, req *pb.RegisterNurseRequest) (*pb.RegisterNurseResponse, error) {
	var err error
	ans := &pb.RegisterNurseResponse{}

	layout := "2006-01-02T15:04:05.000Z"
	dob, err := time.Parse(layout, req.Dob)
	print(dob.String())
	if err != nil {
		fmt.Println(err)
	}

	if len(strings.TrimSpace(req.FirstName)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "First Name cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.LastName)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "Last Name cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.Nic)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "NIC cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.TellNo)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "Tell no cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserEmail)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "Email cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserName)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "UserName cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserPwd)) == 0 {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "Password cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}

	newId := uuid.New()
	id := newId.String()

	userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNICAPI(req.Nic)
	fmt.Println(userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd)
	if err == nil {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "NIC number already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	userTellNo, err := grpca.userapi.GetUserByTellNoAPI(req.TellNo)
	fmt.Println(userTellNo)

	if err == nil {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "TellNo  already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	user_Email, err := grpca.userapi.GetUserByEmailAPI(req.UserEmail)
	fmt.Println(user_Email)
	if err == nil {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "User email already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	user_Id, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNameAPI(req.UserName)
	fmt.Println(user_Id, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd)
	fmt.Println(user_Email)
	if err == nil {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "UserName already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	nurseRegNo, err := grpca.nurseapi.GetUserRegNumberAPI(req.RegNumber)
	fmt.Println(nurseRegNo)
	if err == nil {
		ans.NurseId = 0
		ans.NurseRank = ""
		ans.RegNumber = ""
		ans.Message = "RegNo already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	nurse, err := grpca.userapi.RegisterUserAPI(id, req.UserName, req.FirstName, req.LastName, req.Nic, req.TellNo, req.Address, req.UserEmail, req.UserPwd, req.IsEmailVerified, req.IsTellNoVerified)
	nurseId, regNumber, nurseRank, userId, err := grpca.nurseapi.RegisterNurseAPI(req.RegNumber, dob, req.NurseRank, req.ActiveStatus, req.CenterId, nurse)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error from api")
	}

	ans = &pb.RegisterNurseResponse{
		Status:    http.StatusOK,
		Message:   "Account Created",
		NurseId:   nurseId,
		RegNumber: regNumber,
		NurseRank: nurseRank,
		UserId:    userId,
	}

	return ans, nil
}

func (grpca GRPCAdapter) GetAllNurses(ctx context.Context, req *pb.GetAllNursesRequest) (*pb.GetAllNursesResponse, error) {
	var err error
	ans := &pb.GetAllNursesResponse{}

	if req.CenterId == 0 {
		fmt.Println("GetAllNurses rpc error : center ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	nurses, err := grpca.nurseapi.GetAllNursesAPI(req.CenterId)

	if err != nil {
		fmt.Printf("error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	var nursesList []*pb.Nurse

	for i := 0; i < len(nurses); i++ {
		var tempNurses = &pb.Nurse{
			NurseId:      nurses[i].NurseID,
			RegNumber:    nurses[i].RegNumber,
			Dob:          nurses[i].Dob.String(),
			NurseRank:    nurses[i].NurseRank,
			ActiveStatus: nurses[i].ActiveStatus,
			CenterId:     nurses[i].CenterID,
			UserId:       nurses[i].UserID,
			UserName:     nurses[i].UserName,
			FirstName:    nurses[i].FirstName,
			LastName:     nurses[i].LastName,
			Nic:          nurses[i].Nic,
			TellNo:       nurses[i].TellNo,
			Address:      nurses[i].Address.String,
			UserEmail:    nurses[i].UserEmail,
		}

		nursesList = append(nursesList, tempNurses)
	}

	ans = &pb.GetAllNursesResponse{
		Nurses: nursesList,
	}

	return ans, nil
}

func (grpca GRPCAdapter) GetNurseFilter(ctx context.Context, req *pb.GetNurseFilterRequest) (*pb.GetAllNursesResponse, error) {
	var err error
	ans := &pb.GetAllNursesResponse{}

	if req.CenterId == 0 {
		fmt.Println("Get All Nurses rpc error: center ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required argument")
	}
	var params = db.GetNurseFilterParams{
		CenterID: req.CenterId,
		Column2:  sql.NullString{String: req.Value, Valid: true},
	}

	nurses, err := grpca.nurseapi.GetNurseFilterAPI(params)

	if err != nil {
		fmt.Printf("error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	var nursesList []*pb.Nurse

	for i := 0; i < len(nurses); i++ {
		var tempNurses = &pb.Nurse{
			NurseId:      nurses[i].NurseID,
			RegNumber:    nurses[i].RegNumber,
			Dob:          nurses[i].Dob.String(),
			NurseRank:    nurses[i].NurseRank,
			ActiveStatus: nurses[i].ActiveStatus,
			CenterId:     nurses[i].CenterID,
			UserId:       nurses[i].UserID,
			UserName:     nurses[i].UserName,
			FirstName:    nurses[i].FirstName,
			LastName:     nurses[i].LastName,
			Nic:          nurses[i].Nic,
			TellNo:       nurses[i].TellNo,
			Address:      nurses[i].Address.String,
			UserEmail:    nurses[i].UserEmail,
		}

		nursesList = append(nursesList, tempNurses)
	}

	ans = &pb.GetAllNursesResponse{
		Nurses: nursesList,
	}

	return ans, nil
}

func (grpca GRPCAdapter) UpdateNurse(ctx context.Context, req *pb.UpdateNurseRequest) (*pb.UpdateNurseResponse, error) {
	var err error
	ans := &pb.UpdateNurseResponse{}
	params := db.UpdateUserParams{
		UserID:           req.UserId,
		UserName:         req.UserName,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Nic:              req.Nic,
		TellNo:           req.TellNo,
		Address:          sql.NullString{String: req.Address, Valid: true},
		UserEmail:        req.UserEmail,
		UserPwd:          req.UserPwd,
		IsEmailVerified:  req.IsEmailVerified,
		IsTellNoVerified: req.IsTellNoVerified,
	}

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.Dob)
	nurseParams := db.UpdateNurseParams{
		NurseID:      req.NurseId,
		RegNumber:    req.RegNumber,
		Dob:          time,
		NurseRank:    req.NurseRank,
		ActiveStatus: req.ActiveStatus,
		CenterID:     req.CenterId,
	}

	if err != nil {
		fmt.Println("ubdateNurse rpc : ", err)
	}
	if len(strings.TrimSpace(req.FirstName)) == 0 {

		ans.Message = "First Name cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.LastName)) == 0 {

		ans.Message = "Last Name cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.Nic)) == 0 {

		ans.Message = "NIC cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.TellNo)) == 0 {

		ans.Message = "Tell no cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserEmail)) == 0 {

		ans.Message = "Email cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserName)) == 0 {

		ans.Message = "UserName cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserPwd)) == 0 {

		ans.Message = "Password cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}

	user, err := grpca.userapi.UpdateUserAPI(params)
	nurse, err := grpca.nurseapi.UpdateNurseAPI(nurseParams)
	fmt.Println(user)
	fmt.Println(nurse)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error from api")
	}

	ans = &pb.UpdateNurseResponse{
		Status:  http.StatusOK,
		Message: "Successfully updated",
	}

	return ans, nil

}
