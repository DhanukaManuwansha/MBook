-- name: AdmitPatient :one
INSERT INTO "Patient_Center" ("patient_id", "center_id", "created_at")
VALUES ($1, $2, $3)
RETURNING "patient_id";
