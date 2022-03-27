package ginrestaurantlike

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/restaurantlike/biz"
	restaurantlikemodel "elements-service/modules/restaurantlike/model"
	restaurantlikestorage "elements-service/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{RestaurantId: int(uid.GetLocalID())}

		var paging common.Paging
		err = context.ShouldBind(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewListUsersLikeRestaurantBiz(store)

		users, err := biz.ListUsers(context.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range users {
			users[i].Mask(false)
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
