package elementbiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

type ListElementStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{}, // filter from backend?
		filter *elementmodel.Filter,       // filter from frontend
		paging *common.Paging,
		moreKeys ...string,
	) ([]elementmodel.Element, error)
}

type listElementBiz struct {
	store ListElementStore
}

func NewListElementBiz(store ListElementStore) *listElementBiz {
	return &listElementBiz{store: store}
}

func (biz *listElementBiz) ListElement(
	ctx context.Context,
	recipeId int,
	filter *elementmodel.Filter,
	paging *common.Paging,
) ([]elementmodel.Element, error) {
	result, err := biz.store.ListDataByCondition(ctx, map[string]interface{}{"recipe_id": recipeId}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(elementmodel.EntityName, err)
	}

	return result, nil
}
