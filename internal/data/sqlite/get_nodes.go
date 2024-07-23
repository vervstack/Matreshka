package sqlite

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
)

func (p *Provider) GetNodes(ctx context.Context, names ...string) ([]*evon.Node, error) {
	if len(names) == 0 {
		return nil, nil
	}

	sb := strings.Builder{}
	sb.WriteString(`WHERE `)
	args := make([]any, 0, len(names))
	for i := range names {
		sb.WriteString(`c.name = $`)
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString("\n")
		args = append(args, names[i])

		if i != len(names)-1 {
			sb.WriteString(" OR ")
		}
	}

	row, err := p.conn.QueryContext(ctx, fmt.Sprintf(`
			SELECT 
				cv.key,
				cv.value
			FROM 		configs_values AS cv
			INNER JOIN 	configs 	   AS c  ON c.id = cv.config_id
`+sb.String(), args...))
	if err != nil {
		return nil, errors.Wrap(err, "error getting config values")
	}

	defer row.Close()

	nodes := make([]*evon.Node, 0, len(names))

	for row.Next() {
		node := &evon.Node{}
		err = row.Scan(&node.Name, &node.Value)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning node")
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
