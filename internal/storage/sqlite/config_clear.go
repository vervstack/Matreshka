package sqlite

import (
	"context"

	"go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) ClearValues(ctx context.Context, req domain.ConfigName, version *string) error {
	if version == nil {
		version = toolbox.ToPtr("%%")
	}

	_, err := p.conn.ExecContext(ctx, `
		DELETE FROM configs_values 
	    WHERE config_id = (SELECT id FROM configs WHERE name = $1) 
	    AND   version   LIKE $2`, req.Name(), version)
	if err != nil {
		return rerrors.Wrap(err, "error removing values from configs")
	}

	return nil
}
