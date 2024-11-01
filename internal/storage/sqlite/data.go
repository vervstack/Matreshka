package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/godverv/matreshka-be/internal/clients/sqldb"
	"github.com/godverv/matreshka-be/internal/storage"
)

type Provider struct {
	conn sqldb.DB
}

func New(conn sqldb.DB) *Provider {
	return &Provider{
		conn: conn,
	}
}

func (p *Provider) WithTx(tx *sql.Tx) storage.Data {
	return New(tx)
}
