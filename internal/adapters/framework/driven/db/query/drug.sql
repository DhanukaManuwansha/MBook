-- name: GetDrugById :one
SELECT * FROM "Drug" 
WHERE "drug_id" = $1 
LIMIT 1;
