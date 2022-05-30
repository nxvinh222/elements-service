package recipemodel

const AttributeNameEntityName = "AttributeName"

type AttributeName struct {
	RecipeId int    `json:"recipe_id"`
	Name    string `json:"name"`
}

type AttributeNameCreate struct {
	RecipeId int    `json:"recipe_id"`
	Name    string `json:"name"`
}

func (AttributeName) TableName() string {
	return "attribute_names"
}

func (AttributeNameCreate) TableName() string {
	return AttributeName{}.TableName()
}
