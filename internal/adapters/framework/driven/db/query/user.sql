-- name: RegisterUser :one
INSERT INTO "User" ("user_id", "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "user_pwd", "is_email_verified", "is_tell_no_verified")
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING "user_id";

-- name: GetUserByName :one
SELECT *
FROM "User" 
WHERE "user_name" = $1 
ORDER BY "user_name"
LIMIT 1;

-- name: GetUserByNIC :one
SELECT *
FROM "User"
WHERE "nic" = $1
ORDER BY "nic"
LIMIT 1;

-- name: GetUserByTellNo :one
SELECT *
FROM "User"
WHERE "tell_no" = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM "User"
WHERE "user_email" = $1
LIMIT 1;

-- name: GetAllUsers :many
SELECT *
FROM "User"
ORDER BY "first_name";


-- name: GetAllUsersByNames :many
SELECT *
FROM "User" WHERE "User".user_name LIKE $1||'%' OR  "User".first_name LIKE $1||'%' OR  "User".last_name LIKE $1||'%';




-- name: UpdateUser :one
UPDATE "User"
SET "user_name" = $2, "first_name" = $3 ,"last_name" = $4, "nic" = $5 , "tell_no" =$6 , "address" =$7 , "user_email" =$8 , "user_pwd" =$9, "is_email_verified" =$10, "is_tell_no_verified"=$11
WHERE "user_id" = $1
RETURNING *;
