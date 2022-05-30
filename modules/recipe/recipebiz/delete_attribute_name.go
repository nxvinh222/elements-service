package recipebiz

import "context"

type DeleteAttributeNameStorage interface {
	DeleteAttributeName(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type DeleteAttributeNameBiz struct {
	store DeleteAttributeNameStorage
}

func NewDeleteAttributeNameBiz(store DeleteAttributeNameStorage) *DeleteAttributeNameBiz {
	return &DeleteAttributeNameBiz{store: store}
}

func (biz *DeleteAttributeNameBiz) DeleteAttributeName(ctx context.Context, id int) error {
	err := biz.store.DeleteAttributeName(ctx, map[string]interface{}{"recipe_id": id})
	if err != nil {
		return err
	}

	return nil
}
