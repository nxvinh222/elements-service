package elementstorage

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

func (s *sqlStore) DeleteElement(ctx context.Context, id int) error {
	db := s.db

	err := db.Where("id = ?", id).Delete(elementmodel.Element{}).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
