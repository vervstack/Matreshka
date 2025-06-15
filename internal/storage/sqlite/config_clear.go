package sqlite

import (
	"context"

	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) ClearValues(ctx context.Context, req domain.ConfigName, version string) error {
	_, err := p.conn.ExecContext(ctx, `
		DELETE FROM configs_values 
	    WHERE config_id = (SELECT id FROM configs WHERE name = $1) 
	    AND   version   = $2`, req.Name(), version)
	if err != nil {
		return rerrors.Wrap(err)
	}

	return nil
}
