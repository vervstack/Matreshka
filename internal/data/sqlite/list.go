package sqlite

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) ListConfigs(ctx context.Context, req domain.ListConfigsRequest) ([]string, error) {
	rows, err := p.conn.QueryContext(ctx, `
		SELECT 
		    name
		FROM configs 
		WHERE name LIKE '%'+$1+'%'
		LIMIT $2
		OFFSET $3
		`, req.SearchPattern, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "error listing configs")
	}
	defer rows.Close()

	out := make([]string, 0, req.Limit)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning row")
		}

		out = append(out, name)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error during scanning")
	}

	return out, nil
}
