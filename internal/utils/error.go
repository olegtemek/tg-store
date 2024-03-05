package utils

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type WrappError struct {
	Err error
}

func (re *WrappError) Error() string {
	var pgErr *pgconn.PgError
	if re.Err == pgx.ErrNoRows {
		return "not found"
	}
	if errors.As(re.Err, &pgErr) {
		if pgErr.Code == "23505" {
			return fmt.Sprintf("%s need be unique", pgErr.ConstraintName)
		}
	}

	return fmt.Sprintf("something went wrong %s", re.Err)
}
