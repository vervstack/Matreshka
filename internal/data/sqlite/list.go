package sqlite

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) ListConfigs(ctx context.Context, req domain.ListConfigsRequest) ([]domain.ConfigListItem, error) {
	rows, err := p.conn.QueryContext(ctx, `
		SELECT 
		    cfg.name,
		    version.value
		FROM configs cfg
		LEFT JOIN configs_values AS version
		ON        version.config_id = cfg.id
		AND       version.key       = 'APP-INFO_VERSION'
		WHERE name LIKE '%'||$1||'%'
		LIMIT $2
		OFFSET $3
		`, req.SearchPattern, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "error listing configs")
	}
	defer rows.Close()

	out := make([]domain.ConfigListItem, 0, req.Limit)
	for rows.Next() {
		var item domain.ConfigListItem
		err = rows.Scan(
			&item.Name,
			&item.Version,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning row")
		}

		out = append(out, item)
	}

	return out, nil
}
