package sqlite

import (
	"context"
	"fmt"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka-be/internal/domain"
	api "go.vervstack.ru/matreshka-be/pkg/matreshka_be_api"
)

func (p *Provider) ListConfigs(ctx context.Context, req domain.ListConfigsRequest) (out domain.ListConfigsResponse, err error) {
	err = p.conn.QueryRow(`
			SELECT 
				count(cfg.id)
			FROM configs cfg
			WHERE name LIKE '%'||$1||'%'`, req.SearchPattern).
		Scan(&out.TotalRecords)
	if err != nil {
		return domain.ListConfigsResponse{}, errors.Wrap(err, "error scanning total amount of configs")
	}

	q := `
		SELECT 
		    cfg.name,
		    coalesce(version.value, ''),
		    updated_at
		FROM configs cfg
		LEFT JOIN configs_values AS version
		ON        version.config_id = cfg.id
		AND       version.key       = 'APP-INFO_VERSION'
		WHERE name LIKE '%'||$1||'%'`
	args := []any{req.SearchPattern}

	q += "\nORDER BY " + extractSort(req.Sort)
	q += fmt.Sprintf("\nLIMIT %d OFFSET %d",
		req.Paging.Limit, req.Paging.Offset)

	rows, err := p.conn.QueryContext(ctx, q, args...)
	if err != nil {
		return domain.ListConfigsResponse{}, errors.Wrap(err, "error listing configs")
	}
	defer rows.Close()

	out.List = make([]domain.ConfigListItem, 0, req.Paging.Limit)

	for rows.Next() {
		var item domain.ConfigListItem
		err = rows.Scan(
			&item.Name,
			&item.Version,
			&item.UpdatedAt,
		)
		if err != nil {
			return out, errors.Wrap(err, "error scanning row")
		}

		out.List = append(out.List, item)
	}

	return out, nil
}

func extractSort(sort domain.Sort) (field string) {
	switch sort.SortType {
	case api.Sort_default:
		field = "id "
	case api.Sort_by_name:
		field = "name "
	case api.Sort_by_updated_at:
		field = "updated_at "
	}
	if sort.Desc {
		field += "DESC"
	}

	return
}
