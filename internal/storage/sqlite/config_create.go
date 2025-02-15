package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"
)

func (p *Provider) Create(ctx context.Context, serviceName string) (int64, error) {
	var configId int64

	err := p.conn.QueryRowContext(ctx,
		`
		INSERT INTO configs
				(name)
		VALUES  ($1)
		ON CONFLICT 
		DO UPDATE SET
			     name = excluded.name
		RETURNING COALESCE(
			id, 
			(SELECT id FROM configs WHERE name = $1 LIMIT 1)
	  	)
`,
		serviceName).
		Scan(&configId)
	if err != nil {
		return 0, errors.Wrap(err, "error upserting config")
	}

	return configId, nil
}
