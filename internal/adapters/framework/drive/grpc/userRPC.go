package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var err error
	ans := &pb.RegisterUserResponse{}

	// layout := "2006-01-02T15:04:05.000Z"
	// time, err := time.Parse(layout, req.CreatedAt)
	// print(time.String())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	newId := uuid.New()
	id := newId.String()

	answer, err := grpca.userapi.RegisterUserAPI(id, req.UserName, req.FirstName, req.LastName, req.Nic, req.TellNo, req.Address, req.UserEmail, req.UserPwd, req.IsEmailVerified, req.IsTellNoVerified)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected Error from api")
	}

	ans = &pb.RegisterUserResponse{
		Value: answer,
	}

	return ans, nil

}

func (grpca GRPCAdapter) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetUserByNameResponse, error) {
	var err error
	ans := &pb.GetUserByNameResponse{}

	if req.UserName == "" {
		return ans, status.Error(codes.InvalidArgument, "Missing requiredarguments")
	}

	userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNameAPI(req.UserName)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &pb.GetUserByNameResponse{
		UserId:    userId,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Nic:       nic,
		TellNo:    tellNo,
		Address:   address,
		UserEmail: userEmail,
		UserPwd:   userPwd,
	}

	return ans, nil

}

func (grpca GRPCAdapter) GetUserByNIC(ctx context.Context, req *pb.GetUserByNICRequest) (*pb.GetUserByNICResponse, error) {
	var err error
	ans := &pb.GetUserByNICResponse{}

	if req.Nic == "" {
		return ans, status.Error(codes.InvalidArgument, "Missing required arguments")
	}

	userId, userName, firstName, lastName, nic, tellNo, address, userEmail, userPwd, err := grpca.userapi.GetUserByNICAPI(req.Nic)

	fmt.Printf("UserId: %v", userId)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &pb.GetUserByNICResponse{
		UserId:    userId,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Nic:       nic,
		TellNo:    tellNo,
		Address:   address,
		UserEmail: userEmail,
		UserPwd:   userPwd,
	}

	return ans, nil

}

func (grpca GRPCAdapter) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	fmt.Println("1")

	var err error
	ans := &pb.GetAllUsersResponse{}

	fmt.Println("2")

	users, err := grpca.userapi.GetAllUsersAPI()

	fmt.Println("3")

	if err != nil {
		fmt.Printf("error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	var userList []*pb.User

	for i := 0; i < len(users); i++ {
		var tempUsers = &pb.User{
			UserId:    users[i].UserID,
			UserName:  users[i].UserName,
			FirstName: users[i].FirstName,
			LastName:  users[i].LastName,
			Nic:       users[i].Nic,
			TellNo:    users[i].TellNo,
			Address:   users[i].Address.String,
			UserEmail: users[i].UserEmail,
			UserPwd:   users[i].UserPwd,
		}

		userList = append(userList, tempUsers)
	}

	ans = &pb.GetAllUsersResponse{
		Users: userList,
	}

	return ans, nil
}

func (grpca GRPCAdapter) GetAllUsersByNames(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetAllUsersResponse, error) {
	fmt.Println("1")

	var err error
	ans := &pb.GetAllUsersResponse{}

	fmt.Println("2")

	users, err := grpca.userapi.GetAllUsersByNamesApi(req.UserName)

	fmt.Println("3")

	if err != nil {
		fmt.Printf("error: %v", err)
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	var userList []*pb.User

	for i := 0; i < len(users); i++ {
		var tempUsers = &pb.User{
			UserId:    users[i].UserID,
			UserName:  users[i].UserName,
			FirstName: users[i].FirstName,
			LastName:  users[i].LastName,
			Nic:       users[i].Nic,
			TellNo:    users[i].TellNo,
			Address:   users[i].Address.String,
			UserEmail: users[i].UserEmail,
			UserPwd:   users[i].UserPwd,
		}

		userList = append(userList, tempUsers)
	}

	ans = &pb.GetAllUsersResponse{
		Users: userList,
	}

	return ans, nil
}
