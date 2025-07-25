package sqlite

import (
	"context"

	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) Delete(ctx context.Context, name domain.ConfigName) error {
	_, err := p.conn.ExecContext(ctx, `
			DELETE FROM configs
			WHERE name = $1`,
		name.Name())
	if err != nil {
		return rerrors.Wrap(err, "error deleting config")
	}

	return nil
}
