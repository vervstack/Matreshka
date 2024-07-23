package sqlite

import (
	"context"
	"strconv"
	"strings"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (p *Provider) DeleteValues(ctx context.Context, req domain.PatchConfigRequest) error {
	if len(req.Batch) == 0 {
		return nil
	}

	sb := strings.Builder{}
	sb.WriteString(`AND (`)

	args := make([]any, 0, len(req.Batch)+1)
	args = append(args, req.ServiceName)

	for i, patch := range req.Batch {
		sb.WriteString(`key like $`)
		sb.WriteString(strconv.Itoa(i + 2))
		sb.WriteString(`||'%'`)
		sb.WriteString("\n")
		args = append(args, patch.FieldName)

		if i != len(req.Batch)-1 {
			sb.WriteString(" OR ")
		} else {
			sb.WriteString(")")
		}
	}

	_, err := p.conn.ExecContext(ctx, `
		DELETE FROM configs_values
		WHERE config_id = (
			SELECT
				id
			FROM configs c
			WHERE c.name = $1
)
`+sb.String(), args...)
	if err != nil {
		return errors.Wrap(err, "error deleting values")
	}
	return nil
}
