package upload

import (
	"blog/pkg/file"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"blog/pkg/util"
	"github.com/imroc/req"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.App.PrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.MD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return setting.Image.SavePath
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.Image.AllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logging.Error(err)
		return false
	}

	return size <= setting.Image.MaxSize
}

func CheckImage(src string) error {
	err := file.IsNotExistMkDir(src)
	if err != nil {
		return err
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return err
	}

	return nil
}

func SmUpload(image io.ReadCloser, name string) (*req.Resp, error) {
	// c, _ := os.Open("4c50eef3bdaf0b4164ce179e576f2b2d.jpg")
	header := req.Header{
		"Authorization": setting.Image.SmToken,
	}
	uploadConfig := req.FileUpload{
		File:      image,
		FieldName: "smfile",
		FileName:  name,
	}
	r, err := req.Post("https://sm.ms/api/v2/upload", uploadConfig, header)
	return r, err
}
