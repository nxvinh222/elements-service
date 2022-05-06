package elementstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

func (s *sqlStore) UpdateData(ctx context.Context,
	id int,
	data *elementmodel.ElementUpdate,
) error {
	db := *s.db

	err := db.Where("id = ?", id).Updates(data).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
