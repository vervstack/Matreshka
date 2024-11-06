package sqlite

import (
	"context"
	"fmt"

	errors "github.com/Red-Sock/trace-errors"
)

func (p *Provider) Create(ctx context.Context, serviceName string) error {

	var configId int
	err = p.conn.QueryRowContext(ctx,
		`
		INSERT INTO configs 
		    	(name)
		VALUES  ($1)
		ON CONFLICT DO NOTHING;
		
		SELECT id FROM configs WHERE name = $1;`,
		serviceName).
		Scan(&configId)
	if err != nil {
		return errors.Wrap(err, "error upserting config")
	}

	// TODO если чото отъебёт - потеряем весь конфиг, нужно от обратного: то что не обновляли - удалить
	_, err = p.conn.ExecContext(ctx, `
		DELETE FROM configs_values
	    WHERE config_id = $1`, configId)

	prep, err := p.conn.PrepareContext(ctx, `
		 INSERT INTO configs_values
		    	(config_id, key, value) 
		 VALUES ($1, $2, $3)
		 ON CONFLICT (config_id, key) 
		 DO UPDATE SET value = excluded.value`)
	if err != nil {
		return errors.Wrap(err, "error preparing")
	}

	for _, v := range values {
		if v.Value == nil {
			continue
		}

		if fmt.Sprint(v.Value) == "" {
			continue
		}

		_, err = prep.ExecContext(ctx, configId, v.Name, fmt.Sprint(v.Value))
		if err != nil {
			return errors.Wrap(err, "error inserting config value")
		}
	}

	return nil
}
