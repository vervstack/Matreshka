package sqlite

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) PatchConfig(ctx context.Context, req domain.PatchConfigRequest) error {
	for _, b := range req.Batch {
		_, err := p.conn.ExecContext(ctx,
			`
		INSERT INTO configs_values 
			    (key, value)
		VALUES 	( $1,    $2)`,
			b.FieldName, b.FieldValue)
		if err != nil {
			return errors.Wrap(err, "error upserting config")
		}
	}

	return nil
}
