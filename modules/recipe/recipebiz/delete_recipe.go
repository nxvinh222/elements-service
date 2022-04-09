package recipebiz

import (
	"context"
)

type DeleteRecipeStorage interface {
	DeleteRecipe(ctx context.Context, id int) error
}

type DeleteRecipeBiz struct {
	store DeleteRecipeStorage
}

func NewDeleteRecipeBiz(store DeleteRecipeStorage) *DeleteRecipeBiz {
	return &DeleteRecipeBiz{store: store}
}

func (biz *DeleteRecipeBiz) DeleteRecipe(ctx context.Context, id int) error {
	err := biz.store.DeleteRecipe(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
