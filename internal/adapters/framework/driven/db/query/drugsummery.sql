-- name: GetAllDrugSummery :many
SELECT * FROM "DrugSummery"
WHERE "patient_id"=$1
ORDER BY "created_at" DESC;

-- name: GetAllDrugSummeryOfADrug :many
SELECT "Drug".drug_name,"DrugOrder".dosage ,"DrugOrder".dose,"DrugOrder".frequency, "DrugOrder".givendate, "DrugOrder".giveuntil,"DrugSummery".drugsummery_id,"DrugSummery".summery_date, "DrugSummery".summery_status,"DrugSummery".first_dose_is_given,"DrugSummery".second_dose_is_given,"DrugSummery".third_dose_is_given,"DrugSummery".fourth_dose_is_given,"DrugSummery".drugorder_id,"DrugSummery".patient_id
FROM "DrugSummery" INNER JOIN "DrugOrder"
on "DrugSummery".drugorder_id = "DrugOrder".drugorder_id
Inner Join "Prescription"
on "Prescription".prescription_id = "DrugOrder".prescription_id
Inner Join "Drug"
on "Drug".drug_id = "DrugOrder".drug_id
WHERE "Prescription".patient_id=$1  AND  "Prescription".active_status=TRUE AND "DrugSummery".summery_date>=$2 AND "DrugSummery".summery_date<=$3
;


-- name: UpdateDrugSummery :one
UPDATE "DrugSummery"
SET "summery_date" = $2, "summery_status" = $3 ,"first_dose_is_given" = $4, "second_dose_is_given" = $5 , "third_dose_is_given" =$6, "fourth_dose_is_given" = $7, "drugorder_id" = $8, "patient_id" = $9
WHERE "drugsummery_id" = $1
RETURNING *;