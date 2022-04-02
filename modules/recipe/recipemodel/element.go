package recipemodel

import "elements-service/common"

type Element struct {
	common.SQLModel
	RecipeId int `json:"recipe_id"`
	Name string `json:"name"`
	Selector string `json:"selector"`
	Type string `json:"type"`
	Multiple bool `json:"multiple"`
}

func (Element) TableName() string{
	return "elements"
}

type ElementCreate struct {
	common.SQLModel
	RecipeId int `json:"-"`
	Name string `json:"name"`
	Selector string `json:"selector"`
	Type string `json:"type"`
	Multiple bool `json:"multiple"`
}

func (ElementCreate) TableName() string{
	return Element{}.TableName()
}