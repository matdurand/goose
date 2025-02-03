package gomigrations

import (
	"context"
	"database/sql"

	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationContext(up009, down009)
}

func up009(ctx context.Context, tx *sql.Tx) error {
	return createTable(tx, "echo")
}

func down009(ctx context.Context, tx *sql.Tx) error {
	return dropTable(tx, "echo")
}
