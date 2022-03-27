package restaurantlikemodel

type Filter struct {
	RestaurantId int `json:"-" gorm:"restaurant_id"`
}
