package register

import (
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigration(
		func(_ *sql.Tx) error { return nil },
		func(_ *sql.Tx) error { return nil },
	)
}
