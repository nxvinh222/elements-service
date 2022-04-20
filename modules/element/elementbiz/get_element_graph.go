package elementbiz

import (
	"context"
	"elements-service/modules/element/elementmodel"
)

type GetElementGraphStore interface {
	FindElementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*elementmodel.Element, error)
}

type getElementGraphBiz struct {
	store GetElementGraphStore
}

func NewGetElementGraphBiz(store GetElementGraphStore) *getElementGraphBiz {
	return &getElementGraphBiz{store: store}
}

func (biz *getElementGraphBiz) GetElementGraph(ctx context.Context, fatherId int) ([]elementmodel.ElementReturn, error) {
	var result []elementmodel.ElementReturn

	for true {
		// find element which has this id
		fatherElement, err := biz.store.FindElementByCondition(ctx, map[string]interface{}{"id": fatherId})
		if err != nil {
			return nil, err
		}
		result = append(result, elementmodel.ElementReturn{
			Id:   fatherElement.Id,
			Name: fatherElement.Name,
		})
		if fatherElement.ElementId == nil{
			break
		}
		fatherId = *fatherElement.ElementId
	}

	return result, nil
}
