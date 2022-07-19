package recipebiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/recipe/recipemodel"
)

type ListRecipeStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{}, // filter from backend?
		filter *recipemodel.Filter, // filter from frontend
		paging *common.Paging,
		moreKeys ...string,
	) ([]recipemodel.Recipe, error)
}

type listRecipeBiz struct {
	store ListRecipeStore
}

func NewListRecipeBiz(store ListRecipeStore) *listRecipeBiz {
	return &listRecipeBiz{store: store}
}

func (biz *listRecipeBiz) ListRecipe(
	ctx context.Context,
	userId int,
	filter *recipemodel.Filter,
	paging *common.Paging,
) ([]recipemodel.Recipe, error) {
	result, err := biz.store.ListDataByCondition(ctx, map[string]interface{}{"user_id": userId}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(recipemodel.EntityName, err)
	}

	return result, nil
}
