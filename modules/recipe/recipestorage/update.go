package recipestorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/recipe/recipemodel"
)

func (s *sqlStore) UpdateData(ctx context.Context,
	id int,
	data *recipemodel.RecipeUpdate,
) error {
	db := *s.db

	err := db.Where("id = ?", id).Updates(data).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
