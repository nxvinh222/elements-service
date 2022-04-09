package ginrecipe

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/recipe/recipebiz"
	"elements-service/modules/recipe/recipestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRecipe(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := recipestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := recipebiz.NewDeleteRecipeBiz(store)
		err = biz.DeleteRecipe(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
