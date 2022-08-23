package ports

import db "MedbookServer/internal/adapters/framework/driven/db/sqlc"

type CenterAdminAPIPorts interface {
	RegisterCenterAdminAPI(user_id string, center_id int64) (int64, string, error)
	GetAllCenterAdminsAPI(center_id int64) ([]db.GetAllCenterAdminsRow, error)
	GetAllCenterAdminsByNameApi(arg db.GetAllCenterAdminsByNameParams) ([]db.GetAllCenterAdminsByNameRow, error)
}
