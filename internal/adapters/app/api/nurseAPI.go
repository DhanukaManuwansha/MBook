package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
	"time"
)

type NurseAPIAdapter struct {
	query *db.Queries
}

func NewNurseAPIAdapter(query *db.Queries) *NurseAPIAdapter {
	return &NurseAPIAdapter{
		query: query,
	}
}

func (apiAdpt NurseAPIAdapter) RegisterNurseAPI(reg_number string, dob time.Time, nurse_rank string, active_status bool, center_id int64, user_id string) (int64, string, string, string, error) {

	params := db.RegisterNurseParams{
		RegNumber:    reg_number,
		Dob:          dob,
		NurseRank:    nurse_rank,
		ActiveStatus: active_status,
		CenterID:     center_id,
		UserID:       user_id,
	}

	nurse, err := apiAdpt.query.RegisterNurse(context.Background(), params)

	if err != nil {
		fmt.Println(err)
		return 0, "", "", "", err
	}

	return nurse.NurseID, nurse.RegNumber, nurse.NurseRank, nurse.UserID, nil
}

func (apiAdpt NurseAPIAdapter) GetUserRegNumberAPI(regNo string) (db.Nurse, error) {

	user, err := apiAdpt.query.GetUserRegNumber(context.Background(), regNo)
	if err != nil {
		return db.Nurse{}, err
	}

	return user, nil
}

func (apiAdpt NurseAPIAdapter) GetAllNursesAPI(center_id int64) ([]db.GetAllNursesRow, error) {

	nurses, err := apiAdpt.query.GetAllNurses(context.Background(), center_id)

	if err != nil {
		fmt.Println(err)
		return []db.GetAllNursesRow{}, err
	}

	return nurses, nil
}

func (apiAdpt NurseAPIAdapter) GetNurseFilterAPI(params db.GetNurseFilterParams) ([]db.GetNurseFilterRow, error) {

	nurses, err := apiAdpt.query.GetNurseFilter(context.Background(), params)

	if err != nil {
		fmt.Println(err)
		return []db.GetNurseFilterRow{}, err
	}

	return nurses, nil
}

func (apiAdpt NurseAPIAdapter) UpdateNurseAPI(arg db.UpdateNurseParams) (db.Nurse, error) {

	nurse, err := apiAdpt.query.UpdateNurse(context.Background(), arg)

	if err != nil {
		fmt.Println("UpdateObservation api error : ", err)
		fmt.Println(err.Error())
		return db.Nurse{}, err
	}

	return nurse, nil

}
