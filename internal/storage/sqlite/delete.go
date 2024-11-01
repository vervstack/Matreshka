package sqlite

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) DeleteValues(ctx context.Context, req domain.PatchConfigRequest) error {
	if len(req.Batch) == 0 {
		return nil
	}

	var cfgId int
	err := p.conn.
		QueryRowContext(ctx, `
			SELECT
				id
			FROM configs c
			WHERE c.name = $1`, req.ServiceName).
		Scan(&cfgId)
	if err != nil {
		return errors.Wrap(err, "error getting service id")
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
		return errors.Wrap(err, "error deleting values")
	}

	for _, patch := range req.Batch {
		_, err = deleteQ.ExecContext(ctx, cfgId, patch.FieldName)
		if err != nil {
			return errors.Wrap(err, "error deleting value from db: "+patch.FieldName)
		}
	}

	return nil
}
