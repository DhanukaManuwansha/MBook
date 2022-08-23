-- name: RegisterNurse :one
INSERT INTO "Nurse" ("reg_number", "dob", "nurse_rank", "active_status", "center_id", "user_id")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "nurse_id", "reg_number", "nurse_rank", "user_id";

-- name: GetAllNurses :many
SELECT "nurse_id", "reg_number", "dob", "nurse_rank", "active_status", "center_id", n.user_id, "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "is_email_verified", "is_tell_no_verified"
FROM "Nurse" AS "n"
	INNER JOIN
	"User" AS "u"
	ON n.user_id = u.user_id
WHERE "center_id" = $1
ORDER BY "nurse_id"
LIMIT 10;

-- name: GetNurseFilter :many
SELECT "nurse_id", "reg_number", "dob", "nurse_rank", "active_status", "center_id", n.user_id, "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "is_email_verified", "is_tell_no_verified" 
FROM "Nurse" AS "n"
	INNER JOIN
	"User" AS "u"
	ON n.user_id = u.user_id
WHERE "center_id" = $1 AND ("u".user_name LIKE $2||'%' OR  "u".first_name LIKE $2||'%' OR  "u".last_name LIKE $2||'%' OR  "n".reg_number LIKE $2||'%')
ORDER BY "nurse_id";

-- name: GetUserRegNumber :one
SELECT *
FROM "Nurse"
WHERE "reg_number" = $1
LIMIT 1;

-- name: UpdateNurse :one
UPDATE "Nurse"
SET "reg_number" = $2, "dob" = $3 ,"nurse_rank" = $4, "active_status" = $5 , "center_id" =$6
WHERE "nurse_id" = $1
RETURNING *;
