package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) UpsertValues(ctx context.Context, cfgName string, batch []domain.PatchConfig) error {
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

	for _, b := range batch {
		_, err := p.conn.ExecContext(ctx, `
			INSERT INTO configs_values 
					(config_id, key, value)
			VALUES 	(       $1,  $2,    $3) 
			ON CONFLICT (config_id, key) 
			DO UPDATE SET value = excluded.value`,
			cfgId, b.FieldName, b.FieldValue)
		if err != nil {
			return errors.Wrap(err, "error upserting config")
		}
	}

	return nil
}
