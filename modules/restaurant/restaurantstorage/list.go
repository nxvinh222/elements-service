package restaurantstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{}, // filter from backend?
	filter *restaurantmodel.Filter, // filter from frontend
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if f := filter; f != nil {
		if f.CityId > 0 {
			db = db.Where("city_id = ?", f.CityId)
		}
	}

	err := db.Count(&paging.Total).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	// Count must execute BEFORE preload
	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if c := paging.FakeCursor; c != "" {
		if uid, err := common.FromBase58(c); err == nil {
			db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	err = db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
