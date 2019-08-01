package upload

import (
	"blog/pkg/file"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"blog/pkg/util"
	"fmt"
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
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
