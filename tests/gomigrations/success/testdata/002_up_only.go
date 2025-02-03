package gomigrations

import (
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigration(up002, nil)
}

func up002(tx *sql.Tx) error {
	return createTable(tx, "bravo")
}
