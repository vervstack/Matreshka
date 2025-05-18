package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) DeleteValues(ctx context.Context, req domain.PatchConfigRequest) error {
	if len(req.Delete) == 0 {
		return nil
	}

	cfgId, err := p.getIdByName(ctx, req.ConfigName.Name())
	if err != nil {
		return errors.Wrap(err)
	}

	deleteQ, err := p.conn.PrepareContext(ctx, `
		DELETE FROM configs_values
		WHERE config_id = $1
		AND (
		    key = $2
		    OR 
		    key like $2 ||'_%'
		    )
		AND version = $3`)
	if err != nil {
		return errors.Wrap(err, "error preparing deleting values query")
	}

	for _, valueName := range req.Delete {
		_, err = deleteQ.ExecContext(ctx, cfgId, valueName, req.ConfigVersion)
		if err != nil {
			return errors.Wrap(err, "error deleting value from db: "+valueName)
		}
	}

	return nil
}
