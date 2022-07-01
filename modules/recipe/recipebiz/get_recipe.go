package recipebiz

import (
	"context"
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
	"elements-service/modules/recipe/recipemodel"
	"elements-service/modules/restaurant/restaurantmodel"
)

type GetRecipeStore interface {
	FindRecipeByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*recipemodel.Recipe, error)
}

type getRecipeBiz struct {
	store GetRecipeStore
}

func NewGetRecipeBiz(store GetRecipeStore) *getRecipeBiz{
	return &getRecipeBiz{store: store}
}

func (biz *getRecipeBiz) GetRecipe(ctx context.Context, id int, filter recipemodel.Filter) (*recipemodel.Recipe, error) {
	if filter.Simple == 1 {
		result, err := biz.store.FindRecipeByCondition(ctx, map[string]interface{}{"id": id}, "Elements.ChildElement", "AttributeNameList")
		if err != nil {
			if err == common.RecordNotFound {
				return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
			}

			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		if result.Status == 0 {
			return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
		}

		// Find root element
		var rootElement elementmodel.Element
		var excludeElementList []string
		for _, v := range result.Elements {
			if v.ElementId == nil {
				rootElement = v
				excludeElementList = append(excludeElementList, rootElement.Name)
			}
		}
		// Exclude Object and Link element
		for i := range rootElement.ChildElement{
			if rootElement.ChildElement[i].Type == "object" || rootElement.ChildElement[i].Type == "link" {
				excludeElementList = append(excludeElementList, rootElement.ChildElement[i].Name)
			}
		}

		// Return all attribute name list
		addAttrLoop: for i := range result.AttributeNameList {
			// Remove object and link attribute
			for _, v := range excludeElementList {
				if result.AttributeNameList[i].Name == v{
					continue addAttrLoop
				}
			}
			result.AttributeNameArr = append(result.AttributeNameArr, result.AttributeNameList[i].Name)
		}

		return result, err
	}

	result, err := biz.store.FindRecipeByCondition(ctx, map[string]interface{}{"id": id}, "Elements.ChildElement.ChildElement.ChildElement.ChildElement", "IdentifierList")

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	// Convert Identifier List into array of string
	result.IdentifierArr = make([]string, len(result.IdentifierList))
	for i := range result.IdentifierList{
		result.IdentifierArr[i] = result.IdentifierList[i].Value
	}

	// ??
	for i := len(result.Elements) - 1; i >= 0; i--{
		if result.Elements[i].ElementId != nil {
			result.Elements = RemoveIndex(result.Elements, i)
		}
	}

	return result, err
}

func RemoveIndex(s []elementmodel.Element, index int) []elementmodel.Element {
	ret := make([]elementmodel.Element, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}