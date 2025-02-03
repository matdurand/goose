package gomigrations

import (
	"github.com/matdurand/goose"
)

func init() {
	goose.AddMigrationNoTx(nil, nil)
}
