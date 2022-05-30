package elementstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/attributename/attributenamemodel"
	"elements-service/modules/element/elementmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateData(ctx context.Context,
	id int,
	data *elementmodel.ElementUpdate,
) error {
	tx := *s.db.Begin()

	err := tx.Where("id = ?", id).Updates(data).Error
	if err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	// Find old Attribute name record
	var oldAttrName attributenamemodel.AttributeName
	err = tx.Where("element_id = ? AND name = ?", data.Id, data.OldName).First(&oldAttrName).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return common.ErrDB(err)
		}
	}

	// Delete old Attribute Name
	err = tx.
		Where("element_id = ? AND name = ?", data.Id, data.OldName).
		Delete(attributenamemodel.AttributeName{}).Error
	if err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	// Create Attribute Name
	attrName := attributenamemodel.AttributeName{
		RecipeId: oldAttrName.RecipeId,
		ElementId: oldAttrName.ElementId,
		Name:     data.Name,
	}
	err = tx.Create(&attrName).Error
	if err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	tx.Commit()
	return nil
}
