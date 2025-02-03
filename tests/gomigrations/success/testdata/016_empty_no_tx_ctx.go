package gomigrations

import (
	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationNoTxContext(nil, nil)
}
