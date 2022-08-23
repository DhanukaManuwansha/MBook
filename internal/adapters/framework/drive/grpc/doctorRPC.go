package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) RegisterDoctor(ctx context.Context, req *pb.RegisterDoctorRequest) (*pb.RegisterDoctorResponse, error) {
	var err error
	ans := &pb.RegisterDoctorResponse{}

	layout := "2006-01-02T15:04:05.000Z"
	dob, err := time.Parse(layout, req.Dob)
	print(dob.String())
	if err != nil {
		fmt.Println(err)
	}

	newId := uuid.New()
	id := newId.String()

	if len(strings.TrimSpace(req.FirstName)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "First Name cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.LastName)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "Last Name cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.Nic)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "NIC cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.TellNo)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "Tell no cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserEmail)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "Email cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserName)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "UserName cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	if len(strings.TrimSpace(req.UserPwd)) == 0 {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "Password cannot be null"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNICAPI(req.Nic)
	fmt.Println(userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd)
	if err == nil {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "NIC number already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	userTellNo, err := grpca.userapi.GetUserByTellNoAPI(req.TellNo)
	fmt.Println(userTellNo)

	if err == nil {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "TellNo  already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	user_Email, err := grpca.userapi.GetUserByEmailAPI(req.UserEmail)
	fmt.Println(user_Email)
	if err == nil {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "User email already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	user_Id, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNameAPI(req.UserName)
	fmt.Println(user_Id, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd)
	fmt.Println(user_Email)
	if err == nil {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "UserName already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	doctorRegNo, err := grpca.doctorapi.GetUserByRegNumberAPI(req.RegNumber)
	fmt.Println(doctorRegNo)
	if err == nil {
		ans.DoctorId = 0
		ans.RegNumber = ""
		ans.UserId = ""
		ans.Message = "RegNo already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}

	doctor, err := grpca.userapi.RegisterUserAPI(id, req.UserName, req.FirstName, req.LastName, req.Nic, req.TellNo, req.Address, req.UserEmail, req.UserPwd, req.IsEmailVerified, req.IsTellNoVerified)
	doctorId, regNumber, userId, err := grpca.doctorapi.RegisterDoctorAPI(req.RegNumber, dob, doctor)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error from api")
	}

	ans = &pb.RegisterDoctorResponse{
		DoctorId:  doctorId,
		RegNumber: regNumber,
		UserId:    userId,
	}

	return ans, nil

}
