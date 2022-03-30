package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jared-paxton/cosmerdle_server/pkg/user"
)

type DatabaseAccessor interface {
	CreateAnonymousUser(user user.NewAnonymousUser) error
}

type databaseAccessor struct {
	db *sql.DB
}

func NewDatabaseAccessor(db *sql.DB) DatabaseAccessor {
	return &databaseAccessor{
		db: db,
	}
}

func (dba *databaseAccessor) CreateAnonymousUser(user user.NewAnonymousUser) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT INTO app_user (user_id, created_on, last_activity) VALUES ($1, $2, $3)`

	result, err := dba.db.ExecContext(ctx, statement, user.UserId, user.CreatedOn, user.LastActivity)
	fmt.Println(result)

	return err
}
