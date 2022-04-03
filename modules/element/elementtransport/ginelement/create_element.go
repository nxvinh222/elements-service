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

func CreateElement(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data elementmodel.ElementCreateList

		uid, err := common.FromBase58(context.Param(("id")))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err = context.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := elementstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := elementbiz.NewCreateElementBiz(store)
		err = biz.CreateElement(context.Request.Context(), int(uid.GetLocalID()), &data)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
