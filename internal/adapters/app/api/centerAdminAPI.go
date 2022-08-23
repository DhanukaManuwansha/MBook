package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type CenterAdminAPIAdapter struct {
	query *db.Queries
}

func NewCenterAdminAPIAdapter(query *db.Queries) *CenterAdminAPIAdapter {
	return &CenterAdminAPIAdapter{
		query: query,
	}
}

func (apiAdpt CenterAdminAPIAdapter) RegisterCenterAdminAPI(user_id string, center_id int64) (int64, string, error) {

	params := *&db.RegisterCenterAdminParams{
		UserID:   user_id,
		CenterID: center_id,
	}

	centerAdmin, err := apiAdpt.query.RegisterCenterAdmin(context.Background(), params)

	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}

	return centerAdmin.CenterAdminID, centerAdmin.UserID, nil
}

func (apiAdpt CenterAdminAPIAdapter) GetAllCenterAdminsAPI(center_id int64) ([]db.GetAllCenterAdminsRow, error) {
	centerAdmins, err := apiAdpt.query.GetAllCenterAdmins(context.Background(), center_id)

	if err != nil {
		return []db.GetAllCenterAdminsRow{}, err
	}

	return centerAdmins, nil
}

func (apiAdpt CenterAdminAPIAdapter) GetAllCenterAdminsByNameApi(args db.GetAllCenterAdminsByNameParams) ([]db.GetAllCenterAdminsByNameRow, error) {
	centerAdmins, err := apiAdpt.query.GetAllCenterAdminsByName(context.Background(), args)

	if err != nil {
		return []db.GetAllCenterAdminsByNameRow{}, err
	}

	return centerAdmins, nil
}
