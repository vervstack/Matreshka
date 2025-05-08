package sqlite

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (p *Provider) Rename(ctx context.Context, oldName, newName string) error {
	_, err := p.conn.ExecContext(ctx, `
		UPDATE configs 
		SET name = $1 WHERE name = $2`,
		newName, oldName)
	if err != nil {
		return errors.Wrap(err, "error executing config rename sql")
	}

	return nil
}

func (p *Provider) RenameValues(ctx context.Context, req domain.PatchConfigRequest) error {
	//TODO implement

	return nil
}
