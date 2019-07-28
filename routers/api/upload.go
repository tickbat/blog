package api

import (
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/upload"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	var err error
	code := e.SUCCESS
	status := http.StatusOK
	data := make(map[string]string)
	defer util.Res(c, status, code, data)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn("FormFile error:", err)
		code = e.INVALID_PARAMS
		status = http.StatusBadRequest
		return
	}
	if image == nil {
		code = e.INVALID_PARAMS
		status = http.StatusBadRequest
		return
	}

	imageName := upload.GetImageName(image.Filename)
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		status = http.StatusBadRequest
		return
	}

	savePath := upload.GetImagePath()
	src := savePath + imageName
	if err := upload.CheckImage(savePath); err != nil {
		logging.Warn(err)
		code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
		status = http.StatusInternalServerError
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
		status = http.StatusInternalServerError
		return
	}

	data["image_url"] = upload.GetImageFullUrl(imageName)
}
