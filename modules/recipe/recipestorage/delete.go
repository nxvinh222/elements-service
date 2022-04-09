package recipestorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/recipe/recipemodel"
)

func (s *sqlStore) DeleteRecipe(ctx context.Context, id int) error {
	db := s.db

	err := db.Where("id = ?", id).Delete(recipemodel.Recipe{}).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
