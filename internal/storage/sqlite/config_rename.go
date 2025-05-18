package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) Rename(ctx context.Context, oldName, newName string) error {
	_, err := p.conn.ExecContext(ctx, `
		UPDATE configs 
		SET name = $1 WHERE name = $2`,
		newName, oldName)
	if err != nil {
		return errors.Wrap(err, "error executing config rename sql")
	}

	return nil
}

func (p *Provider) RenameValues(ctx context.Context, req domain.PatchConfigRequest) error {
	if len(req.RenameTo) == 0 {
		return nil
	}

	cfgId, err := p.getIdByName(ctx, req.ConfigName.Name())
	if err != nil {
		return errors.Wrap(err)
	}

	for _, b := range req.RenameTo {
		_, err := p.conn.ExecContext(ctx, `
			UPDATE configs_values 
			SET key = $1
			WHERE config_id = $2
			AND   key = $3
			AND   version = $4
					`,
			b.NewName, cfgId, b.OldName, req.ConfigVersion)
		if err != nil {
			return errors.Wrap(err, "error upserting config")
		}
	}
	return nil
}
