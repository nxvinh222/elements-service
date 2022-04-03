package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
	DbTypeRecipe = 5
	DbTypeElement = 6
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
