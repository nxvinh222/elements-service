package elementmodel

import (
	"elements-service/common"
	"elements-service/modules/attributename/attributenamemodel"
)

const EntityName = "Element"

type Element struct {
	common.SQLModel
	RecipeId  int    `json:"recipe_id" gorm:"column:recipe_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// id of father element id
	ElementId    *int      `json:"element_id" gorm:"column:element_id;"`
	Name         string    `json:"name"`
	Selector     string    `json:"selector"`
	Type         string    `json:"type"`
	Multiple     bool      `json:"multiple"`
	AttributeNameList attributenamemodel.AttributeName `json:"attribute_name_list" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ChildElement []Element `json:"child_elements" gorm:"preload:true;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

type ElementUpdate struct {
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

func (ElementUpdate) TableName() string {
	return Element{}.TableName()
}

type ElementReturn struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (r *Element) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeElement)
}
