package sqlite

import (
	"context"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
)

func (p *Provider) GetConfigNodes(ctx context.Context, serviceName string, version string) (*evon.Node, error) {
	row, err := p.conn.QueryContext(ctx, `
		SELECT
			cv.key,
			coalesce(topv.value, cv.value)
		FROM configs_values AS cv
		INNER JOIN configs    AS c ON c.id = cv.config_id
		AND 	  cv.version 	  = 'master'
		LEFT JOIN configs_values AS topv ON c.id = topv.config_id
		AND topv.key = cv.key
		AND topv.version = $2
		WHERE c.name = $1
		GROUP BY cv.key
`, serviceName, version)
	if err != nil {
		return nil, errors.Wrap(err, "error getting config values")
	}
	defer row.Close()

	var rootNodes []*evon.Node

	for row.Next() {
		var node evon.Node
		err = row.Scan(&node.Name, &node.Value)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning node")
		}

		rootNodes = append(rootNodes, &node)
	}

	if len(rootNodes) == 0 {
		return nil, nil
	}

	ns := evon.NodesToStorage(nil)
	for _, n := range rootNodes {
		ns.AddNode(n)
	}
	return ns[""], nil
}
