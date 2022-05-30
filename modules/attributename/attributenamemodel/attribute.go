package attributenamemodel

const AttributeNameEntityName = "AttributeName"

type AttributeName struct {
	RecipeId int    `json:"recipe_id"`
	ElementId int `json:"element_id"`
	Name    string `json:"name"`
}

type AttributeNameCreate struct {
	RecipeId int    `json:"recipe_id"`
	Name    string `json:"name"`
	OldName string `json:"old_name" gorm:"-"`
}

func (AttributeName) TableName() string {
	return "attribute_names"
}

func (AttributeNameCreate) TableName() string {
	return AttributeName{}.TableName()
}
