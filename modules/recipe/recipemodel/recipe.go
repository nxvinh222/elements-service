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
	IdentifierAttr string `json:"identifier_attr"`
	IdentifierList []Identifier `json:"identifier_list"`
	Elements []elementmodel.Element `json:"elements" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

type RecipeUpdate struct {
	common.SQLModel
	UserId int `json:"user_id"`
	Name string `json:"name"`
	StartUrl string `json:"start_url"`
	Note string `json:"note"`
	IdentifierAttr string `json:"identifier_attr"`
}

func (RecipeCreate) TableName() string {
	return Recipe{}.TableName()
}

func (RecipeUpdate) TableName() string {
	return Recipe{}.TableName()
}

func (r *Recipe) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRecipe)
}

func (r *RecipeCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRecipe)
}
