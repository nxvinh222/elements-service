package recipebiz

import (
	"context"
	"elements-service/modules/recipe/recipemodel"
)

type CreateRecipeStore interface {
	CreateRecipe(ctx context.Context, data *recipemodel.RecipeCreate) error
}

type createRecipeBiz struct {
	store CreateRecipeStore
}

func NewCreateRecipeBiz(store CreateRecipeStore) *createRecipeBiz {
	return &createRecipeBiz{store: store}
}

func (biz *createRecipeBiz) CreateRecipe(ctx context.Context, data *recipemodel.RecipeCreate) error {
	err := biz.store.CreateRecipe(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
