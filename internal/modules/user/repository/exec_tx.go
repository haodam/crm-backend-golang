package repository

import (
	"context"
	"fmt"
)

// ExecTx executes a function within a database transaction
func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v , rb err: %v ", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
