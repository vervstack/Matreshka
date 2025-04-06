package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) DeleteValues(ctx context.Context, req domain.PatchConfigRequest) error {
	if len(req.Batch) == 0 {
		return nil
	}

	var cfgId int64
	err := p.conn.QueryRowContext(ctx, `
		SELECT id 
		FROM configs
		WHERE name = $1
		LIMIT 1`, req.ServiceName).
		Scan(&cfgId)
	if err != nil {
		return errors.Wrap(err, "error getting config id by name")
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

	for _, patch := range req.Batch {
		_, err = deleteQ.ExecContext(ctx, cfgId, patch.FieldName, req.ConfigVersion)
		if err != nil {
			return errors.Wrap(err, "error deleting value from db: "+patch.FieldName)
		}
	}

	return nil
}
