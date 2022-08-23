package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	patientdb "MedbookServer/internal/adapters/framework/driven/db/sqlc/patientDB_sqlc"
	"context"
	"database/sql"
)

//databse ports
type DBPort interface {
	CloseDbConnection()
	AddObservation(ctx context.Context, arg *db.AddObservationParams) error
	UpdateObservation(ctx context.Context, arg *db.UpdateObservationParams) (interface{}, error)
	GetObservationById(ctx context.Context, observationID int64) (*db.Observation, error)

	// user
	RegisterUser(ctx context.Context, arg *db.RegisterUserParams) (string, error)
	GetUserByName(ctx context.Context, userName string) (*db.User, error)
	GetUserByNIC(ctx context.Context, nic string) (*db.User, error)
	GetAllUsers(ctx context.Context) ([]db.User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*db.User, error)
	GetUserByTellNo(ctx context.Context, tellNo string) (*db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	GetAllUsersByNamesGetAllUsersByNames(ctx context.Context, dollar_1 sql.NullString) ([]*db.User, error)

	// ward
	GetWardsByCenter(ctx context.Context, centerID int64) ([]db.Ward, error)

	// bed
	GetBedsByWard(ctx context.Context, wardID int64) ([]db.GetBedsByWardRow, error)

	// super admin
	RegisterSuperAdmin(ctx context.Context, userID string) (*db.SuperAdmin, error) // GetSuperAdmins(ctx context.Context) ([]db.SuperAdmin, error)

	// center admins
	RegisterCenterAdmin(ctx context.Context, arg *db.RegisterCenterAdminParams) (*db.RegisterCenterAdminRow, error)
	GetAllCenterAdmins(ctx context.Context, centerID int64) ([]db.GetAllCenterAdminsRow, error)
	GetAllCenterAdminsByName(ctx context.Context, arg db.GetAllCenterAdminsByNameParams) ([]*db.GetAllCenterAdminsByNameRow, error)

	// nurse
	RegisterNurse(ctx context.Context, arg *db.RegisterNurseParams) (*db.RegisterNurseRow, error)
	GetAllNurses(ctx context.Context, centerID int64) ([]db.GetAllNursesRow, error)
	GetNurseFilter(ctx context.Context, arg *db.GetNurseFilterParams) ([]db.GetNurseFilterRow, error)
	GetUserRegNumber(ctx context.Context, regNumber string) (db.Nurse, error)

	// doctor
	RegisterDoctor(ctx context.Context, arg *db.RegisterDoctorParams) (*db.RegisterDoctorRow, error)
	GetAllBPSisValues(ctx context.Context, patientID string) ([]float64, error)
	GetAllBPDiaValues(ctx context.Context, patientID string) ([]float64, error)
	GetAllObservations(ctx context.Context, patientID string) ([]*db.Observation, error)
	GetAllPRValues(ctx context.Context, patientID string) ([]float64, error)
	GetAllRRValues(ctx context.Context, patientID string) ([]float64, error)
	GetAllTempValues(ctx context.Context, patientID string) ([]float64, error)
	GetUserByRegNumber(ctx context.Context, regNumber string) (db.Doctor, error)

	GetAllNurseNotes(ctx context.Context, patientID string) ([]*db.NurseNotes, error)
	UpdateNurseNote(ctx context.Context, arg *db.UpdateNurseNoteParams) (interface{}, error)
	AddNurseNoteApi(remark, nursingCare, notesDateTimer, createdAt string, nurseId string, patientId string) (int64, error)
	//AddPatient(ctx context.Context, arg *db.AddPatientParams) error
	// GetPatientBYName(ctx context.Context, userName string) ([]*db.User, error)
	// GetPatientById(ctx context.Context, userID string) (*db.User, error)
	// GetPatients(ctx context.Context) ([]*db.User, error)
	// UpdatePatient(ctx context.Context, arg *db.UpdatePatientParams) (interface{}, error)

	GetLatestNurseNotes(ctx context.Context, patientID string) ([]*db.GetLatestNurseNotesRow, error)

	GetActivePrescriptions(ctx context.Context, patientID string) ([]*db.Prescription, error)
	GetLatestPrescriptions(ctx context.Context, patientID string) ([]*db.Prescription, error)

	GetDrugOrdersByPrescriptionID(ctx context.Context, prescriptionID int64) ([]*db.DrugOrder, error)

	GetDrugById(ctx context.Context, drugID int64) (*db.Drug, error)

	GetDrugChart(ctx context.Context, patientID string) ([]*db.GetDrugChartRow, error)
	GetDrugchartForNurseDesk(ctx context.Context, arg *db.GetDrugchartForNurseDeskParams) ([]*db.GetDrugchartForNurseDeskRow, error)

	GetAllDrugSummery(ctx context.Context, patientID string) ([]*db.DrugSummery, error)
	GetAllDrugSummeryOfADrug(ctx context.Context, arg *db.GetAllDrugSummeryOfADrugParams) ([]*db.GetAllDrugSummeryOfADrugRow, error)
	UpdateDrugSummery(ctx context.Context, arg *db.UpdateDrugSummeryParams) (*db.DrugSummery, error)

	GetPatients(ctx context.Context) ([]patientdb.Patient, error)

	AdmitPatient(patientId string, center_id string, created_at string) (int64, error)
	GetPatientBYNIC(ctx context.Context, nic string) ([]patientdb.Patient, error)
	GetPatientBYName(ctx context.Context, firstName string) ([]patientdb.Patient, error)
	GetPatientBYTellNo(ctx context.Context, tellNo string) ([]patientdb.Patient, error)
	GetPatientById(ctx context.Context, patientID string) (patientdb.Patient, error)

	GetAllCenterIds(ctx context.Context) ([]db.Center, error)
}
