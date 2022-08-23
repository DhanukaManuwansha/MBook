// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: nurse.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getAllNurses = `-- name: GetAllNurses :many
SELECT "nurse_id", "reg_number", "dob", "nurse_rank", "active_status", "center_id", n.user_id, "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "is_email_verified", "is_tell_no_verified"
FROM "Nurse" AS "n"
	INNER JOIN
	"User" AS "u"
	ON n.user_id = u.user_id
WHERE "center_id" = $1
ORDER BY "nurse_id"
LIMIT 10
`

type GetAllNursesRow struct {
	NurseID          int64          `json:"nurse_id"`
	RegNumber        string         `json:"reg_number"`
	Dob              time.Time      `json:"dob"`
	NurseRank        string         `json:"nurse_rank"`
	ActiveStatus     bool           `json:"active_status"`
	CenterID         int64          `json:"center_id"`
	UserID           string         `json:"user_id"`
	UserName         string         `json:"user_name"`
	FirstName        string         `json:"first_name"`
	LastName         string         `json:"last_name"`
	Nic              string         `json:"nic"`
	TellNo           string         `json:"tell_no"`
	Address          sql.NullString `json:"address"`
	UserEmail        string         `json:"user_email"`
	IsEmailVerified  bool           `json:"is_email_verified"`
	IsTellNoVerified bool           `json:"is_tell_no_verified"`
}

func (q *Queries) GetAllNurses(ctx context.Context, centerID int64) ([]GetAllNursesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllNurses, centerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllNursesRow
	for rows.Next() {
		var i GetAllNursesRow
		if err := rows.Scan(
			&i.NurseID,
			&i.RegNumber,
			&i.Dob,
			&i.NurseRank,
			&i.ActiveStatus,
			&i.CenterID,
			&i.UserID,
			&i.UserName,
			&i.FirstName,
			&i.LastName,
			&i.Nic,
			&i.TellNo,
			&i.Address,
			&i.UserEmail,
			&i.IsEmailVerified,
			&i.IsTellNoVerified,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNurseFilter = `-- name: GetNurseFilter :many
SELECT "nurse_id", "reg_number", "dob", "nurse_rank", "active_status", "center_id", n.user_id, "user_name", "first_name", "last_name", "nic", "tell_no", "address", "user_email", "is_email_verified", "is_tell_no_verified" 
FROM "Nurse" AS "n"
	INNER JOIN
	"User" AS "u"
	ON n.user_id = u.user_id
WHERE "center_id" = $1 AND ("u".user_name LIKE $2||'%' OR  "u".first_name LIKE $2||'%' OR  "u".last_name LIKE $2||'%' OR  "n".reg_number LIKE $2||'%')
ORDER BY "nurse_id"
`

type GetNurseFilterParams struct {
	CenterID int64          `json:"center_id"`
	Column2  sql.NullString `json:"column_2"`
}

type GetNurseFilterRow struct {
	NurseID          int64          `json:"nurse_id"`
	RegNumber        string         `json:"reg_number"`
	Dob              time.Time      `json:"dob"`
	NurseRank        string         `json:"nurse_rank"`
	ActiveStatus     bool           `json:"active_status"`
	CenterID         int64          `json:"center_id"`
	UserID           string         `json:"user_id"`
	UserName         string         `json:"user_name"`
	FirstName        string         `json:"first_name"`
	LastName         string         `json:"last_name"`
	Nic              string         `json:"nic"`
	TellNo           string         `json:"tell_no"`
	Address          sql.NullString `json:"address"`
	UserEmail        string         `json:"user_email"`
	IsEmailVerified  bool           `json:"is_email_verified"`
	IsTellNoVerified bool           `json:"is_tell_no_verified"`
}

func (q *Queries) GetNurseFilter(ctx context.Context, arg GetNurseFilterParams) ([]GetNurseFilterRow, error) {
	rows, err := q.db.QueryContext(ctx, getNurseFilter, arg.CenterID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNurseFilterRow
	for rows.Next() {
		var i GetNurseFilterRow
		if err := rows.Scan(
			&i.NurseID,
			&i.RegNumber,
			&i.Dob,
			&i.NurseRank,
			&i.ActiveStatus,
			&i.CenterID,
			&i.UserID,
			&i.UserName,
			&i.FirstName,
			&i.LastName,
			&i.Nic,
			&i.TellNo,
			&i.Address,
			&i.UserEmail,
			&i.IsEmailVerified,
			&i.IsTellNoVerified,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserRegNumber = `-- name: GetUserRegNumber :one
SELECT nurse_id, reg_number, dob, nurse_rank, active_status, center_id, user_id
FROM "Nurse"
WHERE "reg_number" = $1
LIMIT 1
`

func (q *Queries) GetUserRegNumber(ctx context.Context, regNumber string) (Nurse, error) {
	row := q.db.QueryRowContext(ctx, getUserRegNumber, regNumber)
	var i Nurse
	err := row.Scan(
		&i.NurseID,
		&i.RegNumber,
		&i.Dob,
		&i.NurseRank,
		&i.ActiveStatus,
		&i.CenterID,
		&i.UserID,
	)
	return i, err
}

const registerNurse = `-- name: RegisterNurse :one
INSERT INTO "Nurse" ("reg_number", "dob", "nurse_rank", "active_status", "center_id", "user_id")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "nurse_id", "reg_number", "nurse_rank", "user_id"
`

type RegisterNurseParams struct {
	RegNumber    string    `json:"reg_number"`
	Dob          time.Time `json:"dob"`
	NurseRank    string    `json:"nurse_rank"`
	ActiveStatus bool      `json:"active_status"`
	CenterID     int64     `json:"center_id"`
	UserID       string    `json:"user_id"`
}

type RegisterNurseRow struct {
	NurseID   int64  `json:"nurse_id"`
	RegNumber string `json:"reg_number"`
	NurseRank string `json:"nurse_rank"`
	UserID    string `json:"user_id"`
}

func (q *Queries) RegisterNurse(ctx context.Context, arg RegisterNurseParams) (RegisterNurseRow, error) {
	row := q.db.QueryRowContext(ctx, registerNurse,
		arg.RegNumber,
		arg.Dob,
		arg.NurseRank,
		arg.ActiveStatus,
		arg.CenterID,
		arg.UserID,
	)
	var i RegisterNurseRow
	err := row.Scan(
		&i.NurseID,
		&i.RegNumber,
		&i.NurseRank,
		&i.UserID,
	)
	return i, err
}

const updateNurse = `-- name: UpdateNurse :one
UPDATE "Nurse"
SET "reg_number" = $2, "dob" = $3 ,"nurse_rank" = $4, "active_status" = $5 , "center_id" =$6
WHERE "nurse_id" = $1
RETURNING nurse_id, reg_number, dob, nurse_rank, active_status, center_id, user_id
`

type UpdateNurseParams struct {
	NurseID      int64     `json:"nurse_id"`
	RegNumber    string    `json:"reg_number"`
	Dob          time.Time `json:"dob"`
	NurseRank    string    `json:"nurse_rank"`
	ActiveStatus bool      `json:"active_status"`
	CenterID     int64     `json:"center_id"`
}

func (q *Queries) UpdateNurse(ctx context.Context, arg UpdateNurseParams) (Nurse, error) {
	row := q.db.QueryRowContext(ctx, updateNurse,
		arg.NurseID,
		arg.RegNumber,
		arg.Dob,
		arg.NurseRank,
		arg.ActiveStatus,
		arg.CenterID,
	)
	var i Nurse
	err := row.Scan(
		&i.NurseID,
		&i.RegNumber,
		&i.Dob,
		&i.NurseRank,
		&i.ActiveStatus,
		&i.CenterID,
		&i.UserID,
	)
	return i, err
}
