package elementbiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/attributename/attributenamemodel"
	"elements-service/modules/element/elementmodel"
	"fmt"
	"strings"
)

type CreateElementStore interface {
	CreateElementList(ctx context.Context, data *elementmodel.ElementCreateList) error

	FindElementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*elementmodel.Element, error)

	FindAttributeNameByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*attributenamemodel.AttributeName, error)
}

type createElementBiz struct {
	store CreateElementStore
}

func NewCreateElementBiz(store CreateElementStore) *createElementBiz {
	return &createElementBiz{store: store}
}

func (biz *createElementBiz) CreateElement(ctx context.Context, recipeId int, data *elementmodel.ElementCreateList) error {
	for i := range data.Elements{
		fatherElement, err := biz.store.FindElementByCondition(ctx, map[string]interface{}{"id": data.Elements[i].ElementId})

		if data.Elements[i].ElementId != nil && err != nil {
			return err
		}

		if fatherElement != nil && fatherElement.Status == 0 {
			return fmt.Errorf("father recipe deleted")
		}

		data.Elements[i].RecipeId = recipeId
		data.Elements[i].Type = strings.ToLower(data.Elements[i].Type)

		// Check if attribute name exists
		attrName, err := biz.store.FindAttributeNameByCondition(ctx, map[string]interface{}{"recipe_id": recipeId, "name": data.Elements[i].Name})
		if err != nil && err != common.RecordNotFound {
			return err
		}
		if attrName != nil {
			return common.ErrEntityExisted(attributenamemodel.AttributeNameEntityName, err)
		}
	}

	err := biz.store.CreateElementList(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
