package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserAPIAdapter struct {
	query *db.Queries
}

func NewUserAPIAdapter(query *db.Queries) *UserAPIAdapter {
	return &UserAPIAdapter{
		query: query,
	}
}

func (apiAdpt UserAPIAdapter) RegisterUserAPI(user_id, user_name, first_name, last_name, nic, tell_no, address, user_email, user_pwd string, is_email_verified, is_tell_no_verified bool) (string, error) {
	//theres no special core logic due to a this is a saving observation in database

	params := *&db.RegisterUserParams{
		UserID:           user_id,
		UserName:         user_name,
		FirstName:        first_name,
		LastName:         last_name,
		Nic:              nic,
		TellNo:           tell_no,
		Address:          sql.NullString{String: address, Valid: true},
		UserEmail:        user_email,
		UserPwd:          user_pwd,
		IsEmailVerified:  is_email_verified,
		IsTellNoVerified: is_tell_no_verified,
	}

	user, err := apiAdpt.query.RegisterUser(context.Background(), params)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return user, nil

}

func (apiAdpt UserAPIAdapter) GetUserByNameAPI(user_name string) (string, string, string, string, string, string, string, string, string, error) {

	user, err := apiAdpt.query.GetUserByName(context.Background(), user_name)
	if err != nil {
		return "", "", "", "", "", "", "", "", "", err
	}

	return user.UserID, user.UserName, user.FirstName, user.LastName, user.Nic, user.TellNo, user.Address.String, user.UserEmail, user.UserPwd, nil

}

func (apiAdpt UserAPIAdapter) GetUserByEmailAPI(email string) (db.User, error) {

	user, err := apiAdpt.query.GetUserByEmail(context.Background(), email)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (apiAdpt UserAPIAdapter) GetUserByNICAPI(nic string) (string, string, string, string, string, string, string, string, string, error) {

	user, err := apiAdpt.query.GetUserByNIC(context.Background(), nic)
	if err != nil {
		return "", "", "", "", "", "", "", "", "", err
	}

	return user.UserID, user.UserName, user.FirstName, user.LastName, user.Nic, user.TellNo, user.Address.String, user.UserEmail, user.UserPwd, nil
}

func (apiAdpt UserAPIAdapter) GetUserByTellNoAPI(tellNo string) (db.User, error) {

	user, err := apiAdpt.query.GetUserByTellNo(context.Background(), tellNo)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (apiAdpt UserAPIAdapter) GetAllUsersAPI() ([]db.User, error) {
	users, err := apiAdpt.query.GetAllUsers(context.Background())

	if err != nil {
		return []db.User{}, err
	}

	return users, nil
}

func (apiAdpt UserAPIAdapter) GetAllUsersByNamesApi(userName string) ([]db.User, error) {
	users, err := apiAdpt.query.GetAllUsersByNames(context.Background(), sql.NullString{String: userName, Valid: true})

	if err != nil {
		return []db.User{}, err
	}

	return users, nil
}

func (apiAdpt UserAPIAdapter) UpdateUserAPI(arg db.UpdateUserParams) (db.User, error) {

	user, err := apiAdpt.query.UpdateUser(context.Background(), arg)

	if err != nil {
		fmt.Println("UpdateObservation api error : ", err)
		fmt.Println(err.Error())
		return db.User{}, err
	}

	return user, nil

}
