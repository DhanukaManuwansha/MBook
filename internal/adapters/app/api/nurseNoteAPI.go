package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type NurseNoteAPIAdapter struct {
	query *db.Queries
}

func NewNurseNoteAPIAdapter(query *db.Queries) *NurseNoteAPIAdapter {
	return &NurseNoteAPIAdapter{
		query: query,
	}
}

//Update nurse notes in API layer					-Dhanuka Manuwansha						-26/05/2022
func (apiAdpt NurseNoteAPIAdapter) UpdateNurseNoteApi(nursingNoteIdIn int64, remarkIn, nursingCareIn string, noteDateTimeIn time.Time, nurseIdIn int64, patientIdIn string) (int64, string, string, time.Time, int64, string, error) {

	params := *&db.UpdateNurseNoteParams{

		NursingnotesID: nursingNoteIdIn,
		Remark:         remarkIn,
		NursingCare:    sql.NullString{String: nursingCareIn, Valid: true},
		NotesDatetime:  sql.NullTime{Time: noteDateTimeIn, Valid: true},
		PatientID:      patientIdIn,
		NurseID:        nurseIdIn,
	}
	nurseNote, err := apiAdpt.query.UpdateNurseNote(context.Background(), params)

	if err != nil {
		fmt.Println("UpdateNurseNote api error : ", err)
		fmt.Println(err.Error())
		return 0, "", "", time.Time{}, 0, "", err
	}

	return nursingNoteIdIn, nurseNote.Remark, nurseNote.NursingCare.String, nurseNote.NotesDatetime.Time, nurseNote.NurseID, nurseNote.PatientID, nil

}

//Get all nurse notes in API layer					-Dhanuka Manuwansha						-26/05/2022
func (apiAdpt NurseNoteAPIAdapter) GetAllNurseNotesApi(patientId string) ([]db.NurseNotes, error) {

	nurseNotes, err := apiAdpt.query.GetAllNurseNotes(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllNurseNotes api error : ", err)
		var errorArray []db.NurseNotes
		return errorArray, err
	}

	return nurseNotes, nil
}

func (apiAdpt NurseNoteAPIAdapter) AddNurseNoteApi(remark, nursingCare string, notesDateTime string, createdAt string, nurseId string, patientId string) (int64, error) {
	//theres no special core logic due to a this is a saving nurse note in database

	//layout := "2006-01-02T15:04:05.000Z"
	//time := time.Now()
	layout := "2006-01-02T15:04:05.000Z"
	newNotesDateTime, err := time.Parse(layout, notesDateTime)

	createdDateTime := time.Now()

	str := nurseId
	n, err := strconv.ParseInt(str, 10, 64)

	params := *&db.AddNurseNoteParams{
		Remark:        remark,
		NursingCare:   sql.NullString{String: nursingCare, Valid: true},
		NotesDatetime: sql.NullTime{Time: newNotesDateTime, Valid: true},
		CreatedAt:     createdDateTime,
		PatientID:     patientId,
		NurseID:       n,
	}

	nursenote, err := apiAdpt.query.AddNurseNote(context.Background(), params)

	if err != nil {
		fmt.Println("AddNurseNote api error : ", err)
		return 0, err
	}

	return nursenote, nil
}

//Get all nurse notes in API layer					-Dhanuka Manuwansha						-07/06/2022
func (apiAdpt NurseNoteAPIAdapter) GetLatestNurseNotesApi(patientId string) ([]db.GetLatestNurseNotesRow, error) {

	nurseNotes, err := apiAdpt.query.GetLatestNurseNotes(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllNurseNotes api error : ", err)
		var errorArray []db.GetLatestNurseNotesRow
		return errorArray, err
	}

	return nurseNotes, nil
}
