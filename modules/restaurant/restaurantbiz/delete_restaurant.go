package restaurantbiz

import (
	"context"
	"fmt"
	"elements-service/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStorage interface {
	SoftDeleteRestaurant(ctx context.Context, id int) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type DeleteRestaurantBiz struct {
	store DeleteRestaurantStorage
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *DeleteRestaurantBiz {
	return &DeleteRestaurantBiz{store: store}
}

func (biz *DeleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return fmt.Errorf("restaurant deleted")
	}

	err = biz.store.SoftDeleteRestaurant(ctx, id)

	return err
}
