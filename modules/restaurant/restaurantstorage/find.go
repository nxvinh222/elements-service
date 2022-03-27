package restaurantstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	err := db.Where(conditions).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, common.RecordNotFound
	}
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
