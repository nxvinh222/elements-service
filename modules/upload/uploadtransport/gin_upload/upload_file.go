package gin_upload

import (
	"fmt"
	"elements-service/common"
	"elements-service/component"
	"elements-service/modules/upload/uploadbusiness"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		_, err = file.Read(dataBytes)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appCtx.GetUploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			fmt.Printf("[ERROR] Cannot upload image %s\n", err)
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
