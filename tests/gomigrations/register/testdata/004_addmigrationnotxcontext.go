package register

import (
	"context"
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationNoTxContext(
		func(_ context.Context, _ *sql.DB) error { return nil },
		func(_ context.Context, _ *sql.DB) error { return nil },
	)
}
