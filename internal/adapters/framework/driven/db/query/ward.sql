-- name: GetWardsByCenter :many
SELECT * FROM "Ward"
WHERE "center_id" = $1
ORDER BY "ward_id";

