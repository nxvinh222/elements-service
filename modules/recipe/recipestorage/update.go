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
	tx := *s.db.Begin()

	// Refresh identifier
	if data.ResetIdentifier == true {
		err := tx.
			Where("recipe_id = ?", data.Id).
			Delete(recipemodel.Identifier{}).Error
		if err != nil {
			tx.Rollback()
			return common.ErrDB(err)
		}
	}

// Update data
	err := tx.Where("id = ?", id).Updates(data).Error
	if err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	tx.Commit()
	return nil
}
