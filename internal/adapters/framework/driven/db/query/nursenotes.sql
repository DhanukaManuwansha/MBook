-- name: GetAllNurseNotes :many
SELECT * FROM "NurseNotes"
WHERE "patient_id"=$1
ORDER BY "nursingnotes_id";

-- name: UpdateNurseNote :one
UPDATE "NurseNotes"
SET "remark" = $2, "nursing_care" = $3, "notes_datetime" = $4 ,  "nurse_id" = $5, "patient_id" = $6
WHERE "nursingnotes_id" = $1
RETURNING *;

-- name: AddNurseNote :one
INSERT INTO "NurseNotes" ( "remark", "nursing_care", "notes_datetime", "created_at", "nurse_id", "patient_id")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "nursingnotes_id" ;
-- name: GetLatestNurseNotes :many
SELECT "NurseNotes".remark, "NurseNotes".nursing_care,"NurseNotes".notes_datetime, "NurseNotes".created_at, "User".user_name
FROM "NurseNotes" INNER JOIN "Nurse"
on "Nurse".nurse_id = "NurseNotes".nurse_id
INNER JOIN "User"
on "Nurse".user_id = "User".user_id
WHERE "patient_id" = $1
ORDER BY "created_at" DESC
LIMIT 14;
