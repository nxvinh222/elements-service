package recipebiz

import (
	"context"
	"elements-service/modules/recipe/recipemodel"
)

type CreateIdentifierListStore interface {
	CreateIdentifierList(ctx context.Context, data []recipemodel.Identifier) error
	DeleteIdentifierListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type createIdentifierListBiz struct {
	store CreateIdentifierListStore
}

func NewCreateIdentifierListBiz(store CreateIdentifierListStore) *createIdentifierListBiz {
	return &createIdentifierListBiz{store: store}
}

func (biz *createIdentifierListBiz) CreateIdentifierList(ctx context.Context, id int, data *recipemodel.IdentifierListCreate) error {
	err := biz.store.DeleteIdentifierListByCondition(ctx, map[string]interface{}{"recipe_id": id})
	if err != nil {
		return err
	}

	if len(data.IdentifierList) == 0 {
		return nil
	}

	for i := range data.IdentifierList {
		data.IdentifierList[i].RecipeId = id
	}

	err = biz.store.CreateIdentifierList(ctx, data.IdentifierList)
	if err != nil {
		return err
	}

	return nil
}
