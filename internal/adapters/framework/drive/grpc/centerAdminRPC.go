package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) RegisterCenterAdmin(ctx context.Context, req *pb.RegisterCenterAdminRequest) (*pb.RegisterCenterAdminResponse, error) {
	var err error
	ans := &pb.RegisterCenterAdminResponse{}

	newId := uuid.New()
	id := newId.String()

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
	userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNICAPI(req.Nic)
	fmt.Println(userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd)
	if err == nil {

		ans.Message = "NIC number already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	userTellNo, err := grpca.userapi.GetUserByTellNoAPI(req.TellNo)
	fmt.Println(userTellNo)

	if err == nil {

		ans.Message = "TellNo  already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	user_Email, err := grpca.userapi.GetUserByEmailAPI(req.UserEmail)
	fmt.Println(user_Email)
	if err == nil {

		ans.Message = "User email already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}
	user_Id, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNameAPI(req.UserName)
	fmt.Println(user_Id, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd)
	fmt.Println(user_Email)
	if err == nil {

		ans.Message = "UserName already added"
		ans.Status = http.StatusBadRequest
		return ans, err
	}

	centerAdmin, err := grpca.userapi.RegisterUserAPI(id, req.UserName, req.FirstName, req.LastName, req.Nic, req.TellNo, req.Address, req.UserEmail, req.UserPwd, req.IsEmailVerified, req.IsTellNoVerified)

	fmt.Print("2")
	centerAdminId, userId, err := grpca.centeradminapi.RegisterCenterAdminAPI(centerAdmin, req.CenterId)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error from api")
	}

	ans = &pb.RegisterCenterAdminResponse{
		CenterAdminId: centerAdminId,
		UserId:        userId,
	}

	return ans, nil

}

func (grpca GRPCAdapter) GetAllCenterAdmins(ctx context.Context, req *pb.GetAllCenterAdminsRequest) (*pb.GetAllCenterAdminsResponse, error) {
	var err error
	ans := &pb.GetAllCenterAdminsResponse{}

	if req.CenterId == 0 {
		fmt.Println("GetAllCenterAdmins rpc error : center ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}

	centerAdmins, err := grpca.centeradminapi.GetAllCenterAdminsAPI(req.CenterId)

	if err != nil {
		fmt.Printf("error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	var centerAdminsList []*pb.CenterAdmin

	for i := 0; i < len(centerAdmins); i++ {
		var tempCenterAminds = &pb.CenterAdmin{
			CenterAdminId: centerAdmins[i].CenterAdminID,
			CenterId:      centerAdmins[i].CenterID,
			UserId:        centerAdmins[i].UserID,
			UserName:      centerAdmins[i].UserName,
			FirstName:     centerAdmins[i].FirstName,
			LastName:      centerAdmins[i].LastName,
			Nic:           centerAdmins[i].Nic,
			TellNo:        centerAdmins[i].TellNo,
			Address:       centerAdmins[i].Address.String,
			UserEmail:     centerAdmins[i].UserEmail,
			UserPwd:       centerAdmins[i].UserPwd,
		}

		centerAdminsList = append(centerAdminsList, tempCenterAminds)
	}

	ans = &pb.GetAllCenterAdminsResponse{
		CenterAdmins: centerAdminsList,
	}

	return ans, nil
}

// INSERT INTO "Center" ("center_id", "center_name", "address", "tell_no", "email", "website", "active_status", "online_status", "created_at", "centertype_id") VALUES (1, "Pvt Hospital", "sri lanka", "1234566788", "hos@email.com", "hos.com", true, true, "2022-06-02 16:52:45+00", 1);
// INSERT INTO CenterType ("centertype_id", "center_type", "centertype_des", "created_at") VALUES (1, "Pvt", "private hospital", "2022-06-02 16:52:45+00", 1);

func (grpca GRPCAdapter) GetAllCenterAdminsByName(ctx context.Context, req *pb.GetAllCenterAdminsByNameRequest) (*pb.GetAllCenterAdminsResponse, error) {
	var err error
	ans := &pb.GetAllCenterAdminsResponse{}

	if req.CenterId == 0 {
		fmt.Println("Get All NurseNotes rpc error : center ID cannot be null")
		return ans, status.Error(codes.InvalidArgument, "missing required params")
	}
	var param = db.GetAllCenterAdminsByNameParams{
		CenterID: req.CenterId,
		Column2:  sql.NullString{String: req.Name, Valid: true},
	}

	centerAdmins, err := grpca.centeradminapi.GetAllCenterAdminsByNameApi(param)

	if err != nil {
		fmt.Printf("error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	var centerAdminsList []*pb.CenterAdmin

	for i := 0; i < len(centerAdmins); i++ {
		var tempCenterAminds = &pb.CenterAdmin{
			CenterAdminId: centerAdmins[i].CenterAdminID,
			CenterId:      centerAdmins[i].CenterID,
			UserId:        centerAdmins[i].UserID,
			UserName:      centerAdmins[i].UserName,
			FirstName:     centerAdmins[i].FirstName,
			LastName:      centerAdmins[i].LastName,
			Nic:           centerAdmins[i].Nic,
			TellNo:        centerAdmins[i].TellNo,
			Address:       centerAdmins[i].Address.String,
			UserEmail:     centerAdmins[i].UserEmail,
			UserPwd:       centerAdmins[i].UserPwd,
		}

		centerAdminsList = append(centerAdminsList, tempCenterAminds)
	}

	ans = &pb.GetAllCenterAdminsResponse{
		CenterAdmins: centerAdminsList,
	}

	return ans, nil
}
