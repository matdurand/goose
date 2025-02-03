package register

import (
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationNoTx(
		func(_ *sql.DB) error { return nil },
		func(_ *sql.DB) error { return nil },
	)
}
