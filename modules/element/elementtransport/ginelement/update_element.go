package ginelement

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/element/elementbiz"
	"elements-service/modules/element/elementmodel"
	"elements-service/modules/element/elementstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateElement(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data elementmodel.ElementUpdate

		id, err := strconv.Atoi(c.Param("id"))

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

		store := elementstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := elementbiz.NewUpdateElementBiz(store)

		err = biz.UpdateRestaurant(c, id, &data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
