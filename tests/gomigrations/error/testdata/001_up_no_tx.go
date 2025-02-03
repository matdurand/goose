package gomigrations

import (
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationNoTx(up001, nil)
}

func up001(db *sql.DB) error {
	q := "CREATE TABLE foo (id INTEGER)"
	_, err := db.Exec(q)
	return err
}
