package elementbiz

import (
	"context"
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
	}

	err := biz.store.CreateElementList(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
