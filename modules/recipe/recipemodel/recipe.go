package recipemodel

import (
	"elements-service/common"
	"elements-service/modules/element/elementmodel"
)

const EntityName = "Recipe"

type Recipe struct {
	common.SQLModel
	UserId int                      `json:"user_id" gorm:"default:0;"`
	Name string                     `json:"name"`
	StartUrl string                 `json:"start_url"`
	Note string                     `json:"note"`
	Elements []elementmodel.Element `json:"elements"`
}

func (Recipe) TableName() string {
	return "recipes"
}

type RecipeCreate struct {
	common.SQLModel
	UserId int `json:"user_id"`
	Name string `json:"name"`
	StartUrl string `json:"start_url"`
	Note string `json:"note"`
}

func (RecipeCreate) TableName() string {
	return Recipe{}.TableName()
}

func (r *RecipeCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}
