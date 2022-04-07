package elementbiz

import (
	"context"
	"elements-service/modules/element/elementmodel"
	"github.com/gin-gonic/gin"
)

type DeleteElementStorage interface {
	DeleteElement(ctx context.Context, id int) error
	FindElementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*elementmodel.Element, error)
}

type DeleteElementBiz struct {
	store DeleteElementStorage
}

func NewDeleteElementBiz(store DeleteElementStorage) *DeleteElementBiz {
	return &DeleteElementBiz{store: store}
}

func (biz *DeleteElementBiz) DeleteElement(ctx context.Context, id int) error {
	_, err := biz.store.FindElementByCondition(ctx, gin.H{"id": id})
	if err != nil {
		return err
	}

	err = biz.store.DeleteElement(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
