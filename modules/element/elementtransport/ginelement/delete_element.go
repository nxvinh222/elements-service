package ginelement

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/element/elementbiz"
	"elements-service/modules/element/elementstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteElement(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := elementstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := elementbiz.NewDeleteElementBiz(store)
		err = biz.DeleteElement(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
