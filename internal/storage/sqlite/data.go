package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"go.vervstack.ru/matreshka-be/internal/clients/sqldb"
	"go.vervstack.ru/matreshka-be/internal/storage"
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
