package restaurantlikestorage

import (
	"context"
	"elements-service/common"
	restaurantlikemodel "elements-service/modules/restaurantlike/model"
)

type likeStore struct {
	RestaurantId int `gorm:"column:restaurant_id;"`
	Count        int `gorm:"column:count;"`
}

func (s *sqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	likeMap := make(map[int]int)

	var likeList []likeStore

	err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&likeList).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	for _, like := range likeList {
		likeMap[like.RestaurantId] = like.Count
	}

	return likeMap, nil
}
