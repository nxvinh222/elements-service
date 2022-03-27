package ginrestaurant

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/restaurant/restaurantbiz"
	"elements-service/modules/restaurant/restaurantmodel"
	"elements-service/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		err := c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewRestaurantBiz(store)
		err = biz.CreateRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
