package storage

import (
	_ "modernc.org/sqlite"

	"github.com/godverv/matreshka-be/internal/clients/sqldb"
)

type Provider struct {
	conn *sqldb.DB
}

func New(conn sqldb.DB) (*Provider, error) {
	conn
	return &Provider{
		conn: conn,
	}, nil
}
