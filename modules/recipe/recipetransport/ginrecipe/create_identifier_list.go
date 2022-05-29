package ginrecipe

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/recipe/recipebiz"
	"elements-service/modules/recipe/recipemodel"
	"elements-service/modules/recipe/recipestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateIdentifierList(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data recipemodel.IdentifierListCreate

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

		store := recipestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := recipebiz.NewCreateIdentifierListBiz(store)

		err = biz.CreateIdentifierList(c, id, &data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
