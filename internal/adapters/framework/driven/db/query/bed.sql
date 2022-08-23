-- name: GetBedsByWard :many
SELECT "bedTicket_id","bed_no","patient_id"  FROM "BedTicket"
WHERE "ward_id" = $1
ORDER BY "bedTicket_id";

-- name: GetPatientsIdByWard :many
SELECT "patient_id" FROM "BedTicket"
WHERE "ward_id" = $1;
