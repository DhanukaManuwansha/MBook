-- name: GetActivePrescriptions :many
SELECT * FROM "Prescription" 
WHERE "patient_id" = $1 And "active_status"= TRUE;


-- name: GetLatestPrescriptions :many
SELECT *
FROM "Prescription" 
WHERE "patient_id" = $1
ORDER BY "created_at" DESC;
