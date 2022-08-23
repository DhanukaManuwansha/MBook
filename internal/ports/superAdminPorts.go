package ports

type SuperAdminAPIPorts interface {
	RegisterSuperAdminAPI(user_id string) (int64, string, error)
	// GetSuperAdmins() ([]db.SuperAdmin, error)
}
