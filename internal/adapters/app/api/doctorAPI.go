package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
	"time"
)

type DoctorAPIAdapter struct {
	query *db.Queries
}

func NewDoctorAPIAdapter(query *db.Queries) *DoctorAPIAdapter {
	return &DoctorAPIAdapter{
		query: query,
	}
}

func (apiAdpt DoctorAPIAdapter) RegisterDoctorAPI(reg_number string, dob time.Time, user_id string) (int64, string, string, error) {

	params := *&db.RegisterDoctorParams{
		RegNumber: reg_number,
		Dob:       dob,
		UserID:    user_id,
	}

	doctor, err := apiAdpt.query.RegisterDoctor(context.Background(), params)

	if err != nil {
		fmt.Println(err)
		return 0, "", "", err
	}

	return doctor.DoctorID, doctor.RegNumber, doctor.UserID, nil
}

func (apiAdpt DoctorAPIAdapter) GetUserByRegNumberAPI(regNumber string) (db.Doctor, error) {

	doc, err := apiAdpt.query.GetUserByRegNumber(context.Background(), regNumber)
	if err != nil {
		return db.Doctor{}, err
	}

	return doc, nil
}
