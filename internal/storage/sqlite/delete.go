package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.verv.tech/matreshka-be/internal/domain"
)

func (p *Provider) DeleteValues(ctx context.Context, cfgName string, batch []domain.PatchConfig) error {
	if len(batch) == 0 {
		return nil
	}

	var cfgId int64
	err := p.conn.QueryRowContext(ctx, `
		SELECT id 
		FROM configs
		WHERE name = $1
		LIMIT 1`, cfgName).
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
		    )`)
	if err != nil {
		return errors.Wrap(err, "error preparing deleting values query")
	}

	for _, patch := range batch {
		_, err = deleteQ.ExecContext(ctx, cfgId, patch.FieldName)
		if err != nil {
			return errors.Wrap(err, "error deleting value from db: "+patch.FieldName)
		}
	}

	return nil
}
