-- name: RegisterSuperAdmin :one
INSERT INTO "SuperAdmin" ("user_id") 
VALUES ($1)
RETURNING "superAdmin_id", "user_id";

-- name: GetSuperAdmins :many
SELECT * FROM "SuperAdmin"
ORDER BY "superAdmin_id";
