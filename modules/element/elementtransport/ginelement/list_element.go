package ginelement

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/element/elementbiz"
	"elements-service/modules/element/elementmodel"
	"elements-service/modules/element/elementstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListElement(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter elementmodel.Filter

		idParam := c.Param("id")
		recipeUid, err := common.FromBase58(idParam)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err = c.ShouldBind(&filter)
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

		store := elementstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := elementbiz.NewListElementBiz(store)
		result, err := biz.ListElement(c.Request.Context(), int(recipeUid.GetLocalID()), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		for i := range result {
			result[i].Mask(false)
			result[i].ElementUid = filter.FatherId
			result[i].ElementId = nil
			result[i].RecipeUid = idParam
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
