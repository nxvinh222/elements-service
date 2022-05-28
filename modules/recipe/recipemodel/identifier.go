package recipemodel

type Identifier struct {
	RecipeId int `json:"recipe_id"`
	Value string `json:"value"`
}

func (Identifier) TableName() string {
	return "identifiers"
}
