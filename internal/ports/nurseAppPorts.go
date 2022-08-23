package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"time"
)

type NurseAPIPorts interface {
	RegisterNurseAPI(reg_number string, dob time.Time, nurse_rank string, active_status bool, center_id int64, user_id string) (int64, string, string, string, error)
	GetAllNursesAPI(center_id int64) ([]db.GetAllNursesRow, error)
	GetNurseFilterAPI(params db.GetNurseFilterParams) ([]db.GetNurseFilterRow, error)
	GetUserRegNumberAPI(regNumber string) (db.Nurse, error)
	UpdateNurseAPI(arg db.UpdateNurseParams) (db.Nurse, error)
}
