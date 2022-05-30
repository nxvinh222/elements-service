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

func CreateAttributeName(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data recipemodel.AttributeNameCreate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err = c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := recipestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := recipebiz.NewCreateAttributeNameBiz(store)

		err = biz.CreateAttributeName(c, id, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
