package elementbiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/attributename/attributenamemodel"
	"elements-service/modules/element/elementmodel"
	"fmt"
)

type UpdateElementStore interface {
	FindElementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*elementmodel.Element, error)

	UpdateData(ctx context.Context,
		id int,
		data *elementmodel.ElementUpdate,
	) error

	FindAttributeNameByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*attributenamemodel.AttributeName, error)
}

type updateElementBiz struct {
	store UpdateElementStore
}

func NewUpdateElementBiz(store UpdateElementStore) *updateElementBiz {
	return &updateElementBiz{store: store}
}

func (biz *updateElementBiz) UpdateElement(ctx context.Context, id int, data *elementmodel.ElementUpdate) error{
	oldData, err := biz.store.FindElementByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return fmt.Errorf("data deleted")
	}

	// Check if attribute name exists
	if oldData.Name != data.Name {
		attrName, err := biz.store.FindAttributeNameByCondition(ctx, map[string]interface{}{"recipe_id": oldData.RecipeId, "name": data.Name})
		if err != nil && err != common.RecordNotFound {
			return err
		}
		if attrName != nil {
			return common.ErrEntityExisted(attributenamemodel.AttributeNameEntityName, err)
		}
	}

	data.Id = id
	data.OldName = oldData.Name
	err = biz.store.UpdateData(ctx, id, data)
	if err != nil {
		return err
	}

	return nil
}
