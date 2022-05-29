package recipebiz

import (
	"context"
	"elements-service/modules/recipe/recipemodel"
	"fmt"
)

type UpdateRecipeStore interface {
	FindRecipeByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*recipemodel.Recipe, error)

	UpdateData(ctx context.Context,
		id int,
		data *recipemodel.RecipeUpdate,
	) error
}

type updateRecipeBiz struct {
	store UpdateRecipeStore
}

func NewUpdateRecipeBiz(store UpdateRecipeStore) *updateRecipeBiz {
	return &updateRecipeBiz{store: store}
}

func (biz *updateRecipeBiz) UpdateRecipe(ctx context.Context, id int, data *recipemodel.RecipeUpdate) error{
	oldData, err := biz.store.FindRecipeByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return fmt.Errorf("data deleted")
	}

	err = biz.store.UpdateData(ctx, id, data)
	if err != nil {
		return err
	}

	return nil
}
