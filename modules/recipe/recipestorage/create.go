package recipestorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/attributename/attributenamemodel"
	"elements-service/modules/recipe/recipemodel"
)

func (s *sqlStore) CreateRecipe(ctx context.Context, data *recipemodel.RecipeCreate) error {
	db := s.db

	err := db.Create(data).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CreateIdentifierList(ctx context.Context, data []recipemodel.Identifier) error {
	db := s.db

	err := db.Create(data).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CreateAttributeName(ctx context.Context, data *attributenamemodel.AttributeNameCreate) error {
	db := s.db

	err := db.Create(data).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
