package elementstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{}, // filter from backend?
	filter *elementmodel.Filter, // filter from frontend
	paging *common.Paging,
	moreKeys ...string,
) ([]elementmodel.Element, error) {
	var result []elementmodel.Element

	db := s.db

	db = db.Table(elementmodel.Element{}.TableName()).Where(conditions).Where("status in (1)")

	if f := filter; f != nil {
		if f.FatherId != "" {
			db = db.Where("element_id = ?", f.FatherIdDecoded)
		} else {
			db = db.Where("element_id IS NULL")
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

	err = db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
