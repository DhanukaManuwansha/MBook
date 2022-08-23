-- name: AddObservation :one
INSERT INTO "Observation" ("obs_datetime", "bp_sys",bp_dia, "pr", "rr", "temp", "notes", "patient_id", "nurse_id") 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING "observation_id" ;

-- name: UpdateObservation :one
UPDATE "Observation"
SET "obs_datetime" = $2, "bp_sys" = $3 ,"bp_dia" = $4, "pr" = $5 , "rr" =$6, "temp" = $7, "notes" = $8, "patient_id" = $9, "nurse_id" =$10
WHERE "observation_id" = $1
RETURNING *;

-- name: GetObservationById :one
SELECT * FROM "Observation" 
WHERE "observation_id" = $1 
LIMIT 1;


-- name: GetAllObservations :many
SELECT * FROM "Observation"
WHERE "patient_id"=$1
ORDER BY "observation_id";

-- name: GetAllBPSisValues :many
SELECT "bp_sys"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY';

-- name: GetAllBPDiaValues :many
SELECT "bp_dia"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY';

-- name: GetAllPRValues :many
SELECT "pr"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY';

-- name: GetAllRRValues :many
SELECT "rr"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY';

-- name: GetAllTempValues :many
SELECT "temp"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY';
