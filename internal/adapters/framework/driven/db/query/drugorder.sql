-- name: GetDrugOrdersByPrescriptionID :many
SELECT * FROM "DrugOrder" 
WHERE "prescription_id" = $1 And "giveuntil">=NOW() AND "givendate"<=Now();
