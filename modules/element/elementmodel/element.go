package elementmodel

import "elements-service/common"

const EntityName = "Element"

type Element struct {
	common.SQLModel
	RecipeId  int    `json:"recipe_id" gorm:"column:recipe_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// id of father element id
	ElementId    *int      `json:"element_id" gorm:"column:element_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name         string    `json:"name"`
	Selector     string    `json:"selector"`
	Type         string    `json:"type"`
	Multiple     bool      `json:"multiple"`
	ChildElement []Element `json:"child_element" gorm:"preload:true"`
}

func (Element) TableName() string {
	return "elements"
}

type ElementCreateList struct {
	Elements []ElementCreate `json:"elements"`
}

type ElementCreate struct {
	common.SQLModel
	RecipeId     int    `json:"-"`
	ElementId    *int   `json:"element_id" gorm:"column:element_id"`
	Name         string `json:"name"`
	Selector     string `json:"selector"`
	Type         string `json:"type"`
	Multiple     bool   `json:"multiple"`
}

func (ElementCreate) TableName() string {
	return Element{}.TableName()
}

func (r *Element) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeElement)
}
