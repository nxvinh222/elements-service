package ginrecipe

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/recipe/recipebiz"
	"elements-service/modules/recipe/recipemodel"
	"elements-service/modules/recipe/recipestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRecipe(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter recipemodel.Filter

		err := c.ShouldBind(&filter)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		err = c.ShouldBind(&paging)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := recipestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := recipebiz.NewListRecipeBiz(store)
		result, err := biz.ListRecipe(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
