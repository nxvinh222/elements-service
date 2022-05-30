package recipebiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/recipe/recipemodel"
)

type CreateAttributeNameStore interface {
	CreateAttributeName(ctx context.Context, data *recipemodel.AttributeNameCreate) error
	FindAttributeNameByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*recipemodel.AttributeName, error)
}

type createAttributeNameBiz struct {
	store CreateAttributeNameStore
}

func NewCreateAttributeNameBiz(store CreateAttributeNameStore) *createAttributeNameBiz {
	return &createAttributeNameBiz{store: store}
}

func (biz *createAttributeNameBiz) CreateAttributeName(ctx context.Context, id int, data *recipemodel.AttributeNameCreate) error {
	attrName, err := biz.store.FindAttributeNameByCondition(ctx, map[string]interface{}{"recipe_id": id, "name": data.Name})
	if err != nil && err != common.RecordNotFound {
		return err
	}
	if attrName != nil {
		return common.ErrEntityExisted(recipemodel.AttributeNameEntityName, err)
	}

	data.RecipeId = id
	err = biz.store.CreateAttributeName(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
