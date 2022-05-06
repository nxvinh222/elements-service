package elementbiz

import (
	"context"
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
}

type updateElementBiz struct {
	store UpdateElementStore
}

func NewUpdateElementBiz(store UpdateElementStore) *updateElementBiz {
	return &updateElementBiz{store: store}
}

func (biz *updateElementBiz) UpdateRestaurant(ctx context.Context, id int, data *elementmodel.ElementUpdate) error{
	oldData, err := biz.store.FindElementByCondition(ctx, map[string]interface{}{"id": id})

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
