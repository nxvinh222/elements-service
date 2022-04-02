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

func CreateRecipe(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data recipemodel.RecipeCreate

		err := context.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := recipestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := recipebiz.NewCreateRecipeBiz(store)
		err = biz.CreateRecipe(context.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
