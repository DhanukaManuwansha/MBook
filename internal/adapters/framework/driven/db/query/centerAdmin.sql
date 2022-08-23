-- name: RegisterCenterAdmin :one
INSERT INTO "CenterAdmin" ("user_id", "center_id")
VALUES ($1, $2)
RETURNING "centerAdmin_id", "user_id";

-- name: GetAllCenterAdmins :many
SELECT "centerAdmin_id", "center_id", ca.user_id, "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "user_pwd", "is_email_verified", "is_tell_no_verified", "created_at"
FROM "CenterAdmin" AS "ca"
    INNER JOIN 
    "User" AS "u"
    ON ca.user_id = u.user_id
WHERE "center_id" = $1
ORDER BY "centerAdmin_id"
LIMIT 10;

-- name: GetAllCenterAdminsByName :many
SELECT "centerAdmin_id", "center_id", ca.user_id, "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "user_pwd", "is_email_verified", "is_tell_no_verified", "created_at"
FROM "CenterAdmin" AS "ca"
    INNER JOIN
    "User" AS "u"
    ON ca.user_id = u.user_id
WHERE "center_id" = $1 AND ("u".user_name LIKE $2||'%' OR  "u".first_name LIKE $2||'%' OR  "u".last_name LIKE $2||'%')
ORDER BY "centerAdmin_id";

