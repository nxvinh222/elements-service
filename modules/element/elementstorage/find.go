package elementstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindElementByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*elementmodel.Element, error) {
	var result elementmodel.Element

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
