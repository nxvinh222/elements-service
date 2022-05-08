package recipemodel

type Filter struct {
	Simple int `json:"simple,omitempty" form:"simple"`
	Name string `json:"name,omitempty" form:"name"`
}
