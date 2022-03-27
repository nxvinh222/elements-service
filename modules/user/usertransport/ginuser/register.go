package ginuser

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/component/hasher"
	"elements-service/modules/user/userbiz"
	"elements-service/modules/user/usermodel"
	"elements-service/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data usermodel.UserCreate

		err := c.ShouldBind(&data)
		if err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		err = biz.Register(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
