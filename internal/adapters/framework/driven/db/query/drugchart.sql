-- name: GetDrugChart :many
SELECT DISTINCT "Drug".drug_id , "Drug".drug_name, "Drug".scientific_name, "Drug".drug_category, "Drug".storage_temperature, "Drug".dangerous_level, "Drug".manufacture , "Drug".no_of_units , "Drug".notes,d.dosage ,d.dose,d.frequency, d.givendate, d.giveuntil
FROM "Drug" INNER JOIN "DrugOrder" AS d
on "Drug".drug_id = d.drug_id 
INNER JOIN "Prescription" AS p
on p.prescription_id = d.prescription_id
WHERE p.patient_id = $1 AND p.active_status =TRUE AND d.giveuntil>=NOW() AND d.givendate<=Now()
ORDER BY "Drug".drug_id;

-- name: GetDrugchartForNurseDesk :many
SELECT "Drug".drug_name,"DrugOrder".dosage ,"DrugOrder".dose,"DrugOrder".frequency, "DrugOrder".givendate, "DrugOrder".giveuntil,"DrugSummery".drugsummery_id,"DrugSummery".summery_date, "DrugSummery".summery_status,"DrugSummery".first_dose_is_given,"DrugSummery".second_dose_is_given,"DrugSummery".third_dose_is_given,"DrugSummery".fourth_dose_is_given,"DrugSummery".drugorder_id,"DrugSummery".patient_id
FROM "DrugSummery" INNER JOIN "DrugOrder"
on "DrugSummery".drugorder_id = "DrugOrder".drugorder_id
Inner Join "Prescription"
on "Prescription".prescription_id = "DrugOrder".prescription_id
Inner Join "Drug"
on "Drug".drug_id = "DrugOrder".drug_id
WHERE "Prescription".patient_id=$1  AND  "Prescription".active_status=TRUE AND "DrugSummery".summery_date=$2  AND "DrugOrder".omit_status = 0;

