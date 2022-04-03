package elementstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

func (s *sqlStore) CreateElementList(ctx context.Context, data *elementmodel.ElementCreateList) error {
	tx := s.db.Begin()

	//tx.Where("element_id = ?", data.Elements[1].ElementId).Delete(&recipemodel.Element{})

	for _, element := range data.Elements{
		err := tx.Create(&element).Error
		if err != nil {
			tx.Rollback()
			return common.ErrDB(err)
		}
	}

	tx.Commit()

	return nil
}
