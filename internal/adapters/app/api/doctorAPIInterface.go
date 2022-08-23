package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"time"
)

type DoctorAPI interface {
	RegisterDoctorAPI(reg_number string, dob time.Time, user_id string) (int64, string, string, error)
	GetUserByRegNumberAPI(regNumber string) (db.Doctor, error)
}
