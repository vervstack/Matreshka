package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Provider struct {
	conn *sql.DB
}

func New(conn *sql.DB) (*Provider, error) {
	return &Provider{
		conn: conn,
	}, nil
}
