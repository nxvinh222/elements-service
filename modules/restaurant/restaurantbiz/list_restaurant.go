package restaurantbiz

import (
	"context"
	"fmt"
	"elements-service/common"
	"elements-service/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore, likeStore LikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, "User") // Keys must have same name with STRUCT ATTRIBUTE NAME aka User

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	likeMap, err := biz.likeStore.GetRestaurantLike(ctx, ids)

	if err != nil {
		fmt.Printf("[ERROR] Cannot get restaurant like: %s\n", err)
	}

	if likeMap != nil {
		for i, id := range ids {
			result[i].LikeCount = likeMap[id]
		}
	}

	return result, nil
}
