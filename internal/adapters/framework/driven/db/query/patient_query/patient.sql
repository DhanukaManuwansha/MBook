-- name: GetPatients :many
SELECT * FROM "Patient" ;

-- name: GetPatientBYName :many
SELECT * FROM "Patient" 
WHERE "first_name" = $1 OR "middle_name" =$1 OR "last_name" = $1;

-- name: GetPatientBYNIC :many
SELECT * FROM "Patient" 
WHERE "nic"= $1;

-- name: GetPatientBYTellNo :many
SELECT * FROM "Patient" 
WHERE "tell_no"= $1;

-- name: GetPatientById :one
SELECT * FROM "Patient" 
WHERE "patient_id"= $1
LIMIT 1;
