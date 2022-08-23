package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
)

type SuperAdminAPIAdapter struct {
	query *db.Queries
}

func NewSuperAdminAPIAdapter(query *db.Queries) *SuperAdminAPIAdapter {
	return &SuperAdminAPIAdapter{
		query: query,
	}
}

func (apiAdpt SuperAdminAPIAdapter) RegisterSuperAdminAPI(user_id string) (int64, string, error) {

	superAdmin, err := apiAdpt.query.RegisterSuperAdmin(context.Background(), user_id)

	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}

	return superAdmin.SuperAdminID, superAdmin.UserID, nil
}

// func (apiAdpt SuperAdminAPIAdapter) GetSuperAdmins() ([]db.SuperAdmin, error) {
// 	superAdmins, err := apiAdpt.query.GetSuperAdmins(context.Background())

// 	if err != nil {
// 		return []db.SuperAdmin{}, err
// 	}

// 	return superAdmins, nil
// }
