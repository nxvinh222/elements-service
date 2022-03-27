package restaurantlikestorage

import (
	"context"
	"fmt"
	"elements-service/common"
	restaurantlikemodel "elements-service/modules/restaurantlike/model"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context,
	conditions map[string]interface{}, // filter from backend?
	filter *restaurantlikemodel.Filter, // filter from frontend
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var results []restaurantlikemodel.Like

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if filter != nil {
		if filter.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", filter.RestaurantId)
		}
	}

	err := db.Count(&paging.Total).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if c := paging.FakeCursor; c != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(c)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	err = db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&results).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(results))
	for i, like := range results {
		users[i] = *like.User
		users[i].CreatedAt = like.CreatedAt
		users[i].UpdatedAt = nil

		if i == len(results)-1 {
			paging.NextCursor = base58.Encode([]byte(fmt.Sprintf("%v", like.CreatedAt.Format(timeLayout))))
		}
	}

	return users, nil
}
