package ports

type AdmitPatientAPIPort interface {
	AdmitPatientApi(patientId string, center_id string, created_at string) (int64, error)
}
