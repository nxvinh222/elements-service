package restaurantbiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/restaurant/restaurantmodel"
)

type GetRestaurantStorage interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type GetRestaurantBiz struct {
	store GetRestaurantStorage
}

func NewGetRestaurantBiz(store GetRestaurantStorage) *GetRestaurantBiz {
	return &GetRestaurantBiz{store: store}
}

func (biz *GetRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return data, err
}
