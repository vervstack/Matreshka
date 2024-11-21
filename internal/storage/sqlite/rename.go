package sqlite

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
)

func (p *Provider) Rename(ctx context.Context, oldName string, newName string) error {
	_, err := p.conn.ExecContext(ctx, `
		UPDATE configs 
		SET name = $1 WHERE name = $2`,
		newName, oldName)
	if err != nil {
		return errors.Wrap(err, "error executing config rename sql")
	}

	return nil
}
