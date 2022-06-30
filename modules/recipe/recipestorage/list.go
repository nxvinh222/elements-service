package recipestorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/recipe/recipemodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{}, // filter from backend?
	filter *recipemodel.Filter, // filter from frontend
	paging *common.Paging,
	moreKeys ...string,
) ([]recipemodel.Recipe, error) {
	var result []recipemodel.Recipe

	db := s.db

	db = db.Table(recipemodel.Recipe{}.TableName()).Where(conditions).Where("status <> 0")

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
