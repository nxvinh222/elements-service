package biz

import (
	"context"
	"elements-service/common"
	restaurantlikemodel "elements-service/modules/restaurantlike/model"
)

type ListUsersLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context,
		conditions map[string]interface{}, // filter from backend?
		filter *restaurantlikemodel.Filter, // filter from frontend
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUsersLikeRestaurantBiz struct {
	store ListUsersLikeRestaurantStore
}

func NewListUsersLikeRestaurantBiz(store ListUsersLikeRestaurantStore) *listUsersLikeRestaurantBiz {
	return &listUsersLikeRestaurantBiz{store: store}
}

func (biz listUsersLikeRestaurantBiz) ListUsers(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
