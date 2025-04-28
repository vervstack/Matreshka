package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
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
		WITH cfg AS (
			SELECT
				configs.id 			AS id,
				configs.updated_at 	AS updated_at,
				configs.name 		AS name,
				cv.version 			AS version
			FROM configs
			JOIN configs_values cv on configs.id = cv.config_id
			WHERE name LIKE '%'||$1||'%'
			GROUP BY configs.name, cv.version
		)
		SELECT
			cfg.name 				as service_name,
			service_version.value 			as service_version,
			cfg.updated_at 					as last_updated_at, 
			json_group_array(cfg.version) 	as config_versions
		FROM cfg
		LEFT JOIN   configs_values AS service_version
		ON          service_version.config_id = cfg.id
		AND         service_version.key       = 'APP-INFO_VERSION'
		AND         service_version.version   = 'master'
		GROUP BY cfg.id
		HAVING COUNT(cfg.id) > 0  -- Ensures only non-empty results are returned
		`
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
		var versionsJSON string
		err = rows.Scan(
			&item.Name,
			&item.Version,
			&item.UpdatedAt,
			&versionsJSON,
		)
		if err != nil {
			return out, errors.Wrap(err, "error scanning row")
		}

		err = json.Unmarshal([]byte(versionsJSON), &item.ConfigVersions)
		if err != nil {
			return out, errors.Wrap(err, "error marshalling from json ")
		}
		sort.Slice(item.ConfigVersions, func(i, j int) bool {
			return item.ConfigVersions[i] < item.ConfigVersions[j]
		})

		for i := range item.ConfigVersions {
			if item.ConfigVersions[i] == domain.MasterVersion {
				item.ConfigVersions[0], item.ConfigVersions[i] =
					item.ConfigVersions[i], item.ConfigVersions[0]

				break
			}
		}

		out.List = append(out.List, item)
	}

	return out, nil
}

func (p *Provider) GetVersions(ctx context.Context, name string) ([]string, error) {
	q := `
		WITH cfg AS (
				SELECT 
				    configs.id         AS id,
                    configs.updated_at AS updated_at,
                    configs.name       AS name,
                    cv.version     AS version
             FROM configs
			 JOIN configs_values cv ON configs.id = cv.config_id
             WHERE name = $1
             GROUP BY configs.name, cv.version)
SELECT json_group_array(cfg.version)
FROM cfg
LEFT JOIN configs_values AS service_version
ON service_version.config_id = cfg.id
AND service_version.key = 'APP-INFO_VERSION'
AND service_version.version = 'master'
GROUP BY cfg.id
HAVING COUNT(cfg.id) > 0 -- Ensures only non-empty results are returned
		`

	var versionsStr []byte
	err := p.conn.QueryRowContext(ctx, q, name).
		Scan(&versionsStr)
	if err != nil {
		return nil, errors.Wrap(err, "error getting versions")
	}

	var versions []string
	err = json.Unmarshal(versionsStr, &versions)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling versions from json ")
	}

	return versions, nil
}

func extractSort(sort domain.Sort) (field string) {
	switch sort.SortType {
	case api.Sort_default:
		field = "id"
	case api.Sort_by_name:
		field = "name"
	default:
		field = "updated_at"
	}
	if sort.Desc {
		field += " DESC"
	}

	return
}
