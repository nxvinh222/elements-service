package recipebiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/attributename/attributenamemodel"
)

type CreateAttributeNameStore interface {
	CreateAttributeName(ctx context.Context, data *attributenamemodel.AttributeNameCreate) error
	FindAttributeNameByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*attributenamemodel.AttributeName, error)
	DeleteAttributeName(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type createAttributeNameBiz struct {
	store CreateAttributeNameStore
}

func NewCreateAttributeNameBiz(store CreateAttributeNameStore) *createAttributeNameBiz {
	return &createAttributeNameBiz{store: store}
}

func (biz *createAttributeNameBiz) CreateAttributeName(ctx context.Context, id int, data *attributenamemodel.AttributeNameCreate) error {
	// Do nothing if new name is the same as old name
	if data.Name == data.OldName {
		return nil
	}

	// If new name and old name different
	// Check if new name exists
	attrName, err := biz.store.FindAttributeNameByCondition(ctx, map[string]interface{}{"recipe_id": id, "name": data.Name})
	if err != nil && err != common.RecordNotFound {
		return err
	}
	if attrName != nil {
		return common.ErrEntityExisted(attributenamemodel.AttributeNameEntityName, err)
	}

	// If new name not exists, delete old name
	err = biz.store.DeleteAttributeName(ctx, map[string]interface{}{"recipe_id": id, "name": data.OldName})
	if err != nil {
		return err
	}

	// Create new name
	data.RecipeId = id
	err = biz.store.CreateAttributeName(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
