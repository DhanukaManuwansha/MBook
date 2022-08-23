package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type userAPI interface {
	RegisterUserAPI(user_id, user_name, firstName, last_name, nic, tell_no, address, user_email, user_pwd string, is_email_verified, is_tell_no_verified bool) (string, error)
	GetUserByNameAPI(user_name string) (string, string, string, string, string, string, string, string, string, error)
	GetUserByNICAPI(nic string) (string, string, string, string, string, string, string, string, string, error)
	GetAllUsersAPI() ([]db.User, error)
	GetAllUsersByNamesApi(user_name string) ([]db.User, error)
	GetUserByEmailAPI(userEmail string) (db.User, error)
	GetUserByTellNoAPI(tellNo string) (db.User, error)
	UpdateUserAPI(arg db.UpdateUserParams) (db.User, error)
}
