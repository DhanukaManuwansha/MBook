// Thaveesha - 30/05/2022
package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
)

type WardAPIAdapter struct {
	query *db.Queries
}

func NewWardAPIAdapter(query *db.Queries) *WardAPIAdapter {
	return &WardAPIAdapter{
		query: query,
	}
}

func (apiAdpt WardAPIAdapter) GetWardsByCenterAPI(center_id int64) ([]db.Ward, error) {

	wards, err := apiAdpt.query.GetWardsByCenter(context.Background(), center_id)

	if err != nil {
		fmt.Printf("Cannot load wards from API layer: %v", err)
		return []db.Ward{}, err
	}

	return wards, nil
}
