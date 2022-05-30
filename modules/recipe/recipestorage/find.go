package recipestorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/recipe/recipemodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindRecipeByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*recipemodel.Recipe, error) {
	var result recipemodel.Recipe

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	err := db.Where(conditions).Find(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, common.RecordNotFound
	}
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (s *sqlStore) FindAttributeNameByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*recipemodel.AttributeName, error) {
	var result recipemodel.AttributeName

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
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
