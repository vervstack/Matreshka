package sqlite

import (
	"context"
	"time"

	errors "go.redsock.ru/rerrors"
)

func (p *Provider) SetUpdatedAt(ctx context.Context, serviceName string, time time.Time) error {
	_, err := p.conn.ExecContext(ctx, `
		UPDATE configs
		SET updated_at = $1
		WHERE name = $2`, time, serviceName)
	if err != nil {
		return errors.Wrap(err, "error updating updated_at for config")
	}

	return nil
}
