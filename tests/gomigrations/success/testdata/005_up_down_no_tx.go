package gomigrations

import (
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationNoTx(up005, down005)
}

func up005(db *sql.DB) error {
	return createTable(db, "charlie")
}

func down005(db *sql.DB) error {
	return dropTable(db, "charlie")
}
