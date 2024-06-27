package sqlite

import (
	"context"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) GetConfig(ctx context.Context, req domain.GetConfigReq) (*evon.Node, error) {
	row, err := p.conn.QueryContext(ctx, `
			SELECT 
				cv.key,
				cv.value
			FROM 		configs_values AS cv
			INNER JOIN 	configs 	   AS c  ON c.id = cv.config_id
			AND 		c.name 				  = $1
`, req.ServiceName)
	if err != nil {
		return nil, errors.Wrap(err, "error getting config values")
	}
	defer row.Close()

	root := &evon.Node{}

	for row.Next() {
		var node evon.Node
		err = row.Scan(&node.Name, &node.Value)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning node")
		}

		root.InnerNodes = append(root.InnerNodes, &node)
	}

	err = row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error during scanning")
	}

	return root, nil
}
