package elementbiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

type GetElementStorage interface {
	FindElementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*elementmodel.Element, error)
}

type GetElementBiz struct {
	store GetElementStorage
}

func NewGetElementBiz(store GetElementStorage) *GetElementBiz {
	return &GetElementBiz{store: store}
}

func (biz *GetElementBiz) GetElement(ctx context.Context, id int) (*elementmodel.Element, error) {
	data, err := biz.store.FindElementByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(elementmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(elementmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(elementmodel.EntityName, nil)
	}

	return data, err
}

