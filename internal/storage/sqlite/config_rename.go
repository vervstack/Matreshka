package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"
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
