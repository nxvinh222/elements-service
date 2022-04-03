package elementmodel

import "elements-service/common"

type Element struct {
	common.SQLModel
	RecipeId int `json:"recipe_id"`
	// id of father element id
	ElementId *int         `json:"element_id"`
	Name string            `json:"name"`
	Selector string        `json:"selector"`
	Type string            `json:"type"`
	Multiple bool          `json:"multiple"`
	ChildElement []Element `json:"child_element"`
}

func (Element) TableName() string{
	return "elements"
}

type ElementCreateList struct {
	Elements []ElementCreate `json:"elements"`
}

type ElementCreate struct {
	common.SQLModel
	RecipeId int `json:"-"`
	ElementIdStr string `json:"element_id" gorm:"-"`
	ElementId *int `gorm:"column:element_id"`
	Name string `json:"name"`
	Selector string `json:"selector"`
	Type string `json:"type"`
	Multiple bool `json:"multiple"`
}

func (ElementCreate) TableName() string{
	return Element{}.TableName()
}