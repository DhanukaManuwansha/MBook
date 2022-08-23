-- name: RegisterDoctor :one
INSERT INTO "Doctor" ("reg_number", "dob", "user_id")
VALUES ($1, $2, $3)
RETURNING "doctor_id", "reg_number", "user_id";

-- name: GetUserByRegNumber :one
SELECT *
FROM "Doctor"
WHERE "reg_number" = $1
LIMIT 1;
