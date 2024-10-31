// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: pre_go_crm_user_c.sql

package database

import (
	"context"
)

const getUserByEmailSQLC = `-- name: GetUserByEmailSQLC :one
SELECT user_id, user_email FROM pre_go_crm_user_c WHERE user_email = ? LIMIT 1
`

type GetUserByEmailSQLCRow struct {
	UserID    uint32
	UserEmail string
}

func (q *Queries) GetUserByEmailSQLC(ctx context.Context, userEmail string) (GetUserByEmailSQLCRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailSQLC, userEmail)
	var i GetUserByEmailSQLCRow
	err := row.Scan(&i.UserID, &i.UserEmail)
	return i, err
}
