package sqlite

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) UpdateValues(ctx context.Context, req domain.PatchConfigRequest) error {
	var configId int
	err := p.conn.QueryRowContext(ctx, `
			SELECT 
				id
			FROM configs 
			WHERE name = $1`,
		req.ServiceName).
		Scan(&configId)
	if err != nil {
		return errors.Wrap(err, "error getting service config id")
	}

	for _, b := range req.Batch {
		_, err := p.conn.ExecContext(ctx, `
			INSERT INTO configs_values 
					(config_id, key, value)
			VALUES 	(       $1,  $2,    $3) 
			ON CONFLICT (config_id, key) 
			DO UPDATE SET value = excluded.value`,
			configId, b.FieldName, b.FieldValue)
		if err != nil {
			return errors.Wrap(err, "error upserting config")
		}
	}

	return nil
}
