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

func UpdateRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		uid, err := common.FromBase58(c.Param(("id")))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid id",
			})

			return
		}

		err = c.ShouldBind(&data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		err = biz.UpdateRestaurant(c, int(uid.GetLocalID()), &data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
