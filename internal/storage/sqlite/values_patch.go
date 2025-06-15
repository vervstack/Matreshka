package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) UpsertValues(ctx context.Context, req domain.PatchConfigRequest) error {
	if len(req.Upsert) == 0 {
		return nil
	}

	cfgId, err := p.getIdByName(ctx, req.ConfigName.Name())
	if err != nil {
		return errors.Wrap(err)
	}

	for _, b := range req.Upsert {
		_, err := p.conn.ExecContext(ctx, `
			INSERT INTO configs_values 
					(config_id, key, value, version)
			VALUES 	(       $1,  $2,    $3,      $4) 
			ON CONFLICT (config_id, key, version) 
			DO UPDATE SET value = excluded.value`,
			cfgId, b.FieldName, b.FieldValue, req.ConfigVersion)
		if err != nil {
			return errors.Wrap(err, "error upserting config")
		}
	}

	return nil
}

func (p *Provider) getIdByName(ctx context.Context, name string) (cfgId int64, err error) {
	err = p.conn.QueryRowContext(ctx, `
		SELECT id 
		FROM configs
		WHERE name = $1
		LIMIT 1`, name).
		Scan(&cfgId)
	if err != nil {
		return 0, errors.Wrap(err, "error getting config id by name")
	}

	return cfgId, nil
}
