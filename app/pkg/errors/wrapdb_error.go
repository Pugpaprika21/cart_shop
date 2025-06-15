package errors

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgconn"
)

func WrapDBError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return fmt.Errorf("MySQL Error [%d]: %s", mysqlErr.Number, mysqlErr.Message)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return fmt.Errorf("PostgreSQL Error [%s]: %s", pgErr.Code, pgErr.Message)
	}

	return fmt.Errorf("Unknown DB Error: %v", err)
}
