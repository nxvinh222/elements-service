package ginuser

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/component/hasher"
	"elements-service/component/tokenprovider/jwt"
	"elements-service/modules/user/userbiz"
	"elements-service/modules/user/usermodel"
	"elements-service/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLoginData usermodel.UserLogin

		err := c.ShouldBind(&userLoginData)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		tProvider := jwt.NewJwtProvider(ctx.GetSecretKey())
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &userLoginData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
