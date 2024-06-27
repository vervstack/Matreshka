package sqlite

import (
	"context"
	"fmt"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
)

func (p *Provider) SaveConfig(ctx context.Context, serviceName string, cfg matreshka.AppConfig) error {
	node, err := evon.MarshalEnv(&cfg)
	if err != nil {
		return errors.Wrap(err, "error marshalling config to variables")
	}
	values := evon.NodesToStorage(node.InnerNodes)

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
	for _, v := range values {
		if v.Value == nil {
			continue
		}

		if fmt.Sprint(v.Value) == "" {
			continue
		}

		_, err = p.conn.ExecContext(ctx, `
		INSERT INTO configs_values
		    	(config_id, key, value) 
		 VALUES ($1, $2, $3)`,
			configId, v.Name, fmt.Sprint(v.Value))
		if err != nil {
			return errors.Wrap(err, "")
		}
	}

	return nil
}
