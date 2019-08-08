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
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Error("FormFile error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if image == nil {
		util.Res(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		util.Res(c, http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	if c.Request.URL.Path == "/sm/upload" {
		r, err := upload.SmUpload(file, imageName)
		if err != nil {
			logging.Error("upload to sm.ms error:", err)
			util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		var res map[string]interface{}
		err = r.ToJSON(&res)
		if err != nil {
			logging.Error("parse json from sm.ms error:", err)
			util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		if res["success"] != true {
			logging.Error("sm.ms response error:", res["message"])
			util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		util.Res(c, http.StatusOK, e.SUCCESS, res["data"])
	} else {
		savePath := upload.GetImagePath()
		src := savePath + imageName
		if err := upload.CheckImage(savePath); err != nil {
			logging.Error("check image error:", err)
			util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}

		if err := c.SaveUploadedFile(image, src); err != nil {
			logging.Error("save upload file error:", err)
			util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		util.Res(c, http.StatusOK, e.SUCCESS, gin.H{
			"image_url": upload.GetImageFullUrl(imageName),
		})
	}
}
