package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type AdmitPatientAPIAdapter struct {
	query *db.Queries
}

func NewAdmitPatientAPIAdapter(query *db.Queries) *AdmitPatientAPIAdapter {
	return &AdmitPatientAPIAdapter{
		query: query,
	}
}

func (apiAdpt AdmitPatientAPIAdapter) AdmitPatientApi(patientId string, center_id string, created_at string) (int64, error) {
	//theres no special core logic due to a this is a saving admitpatient in database

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, created_at)

	str := center_id
	n, err := strconv.ParseInt(str, 10, 64)

	params := *&db.AdmitPatientParams{
		PatientID: patientId,
		CenterID:  n,
		CreatedAt: time,
	}

	admitpatient, err := apiAdpt.query.AdmitPatient(context.Background(), params)
	admitpatientreturn, err := strconv.ParseInt(admitpatient, 10, 64)

	if err != nil {
		fmt.Println("AddObservation api error : ", err)
		return 0, err
	}

	return admitpatientreturn, nil

}
