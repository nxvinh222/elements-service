package recipemodel

type Identifier struct {
	RecipeId int    `json:"recipe_id"`
	Value    string `json:"value"`
}

type IdentifierListCreate struct {
	IdentifierList []Identifier `json:"identifier_list"`
}

func (Identifier) TableName() string {
	return "identifiers"
}
