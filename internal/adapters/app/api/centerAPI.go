package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type CenterAPIAdapter struct {
	query *db.Queries
}

func NewCenterAPIAdapter(query *db.Queries) *CenterAPIAdapter {
	return &CenterAPIAdapter{
		query: query,
	}
}

func (apiAdpt CenterAPIAdapter) GetAllCenterIds() ([]db.Center, error) {

	centers, err := apiAdpt.query.GetAllCenterIds(context.Background())

	if err != nil {
		fmt.Println("GetAllBPValues api error : ", err)

		return []db.Center{}, err
	}

	return centers, nil
}
