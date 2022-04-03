package elementmodel

type Filter struct {
	FatherId        string `json:"father_id,omitempty" form:"father_id"`
	FatherIdDecoded int    `json:"-" form:"-"`
}
