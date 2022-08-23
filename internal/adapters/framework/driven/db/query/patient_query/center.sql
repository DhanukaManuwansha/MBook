-- name: GetAllObservations :many
SELECT "center_id" , "center_name"
 FROM "Center"
ORDER BY "center_id";