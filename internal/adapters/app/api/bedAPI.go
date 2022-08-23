// Thaveesha - 30/05/2022
package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
)

type BedAPIAdapter struct {
	query *db.Queries
}

func NewBedAPIAdapter(query *db.Queries) *BedAPIAdapter {
	return &BedAPIAdapter{
		query: query,
	}
}

func (apiAdpt BedAPIAdapter) GetBedsByWardApi(ward_id int64) ([]db.GetBedsByWardRow, error) {

	beds, err := apiAdpt.query.GetBedsByWard(context.Background(), ward_id)

	if err != nil {
		fmt.Printf("Cannot load beds from API layer: %v", err)
		return []db.GetBedsByWardRow{}, err
	}

	return beds, nil
}
