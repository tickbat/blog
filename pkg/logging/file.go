package logging

import (
	"blog/pkg/setting"
	"fmt"
	"os"
	"time"
	"blog/pkg/file"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.Log.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", setting.Log.LogSaveName, time.Now().Format(setting.Log.TimeFormat), setting.Log.LogFileExt)
}

func openLogFile(name, path string) (*os.File, error) {
	// dir, err := os.Getwd()
    // if err != nil {
    //     return nil, fmt.Errorf("os.Getwd err: %v", err)
    // }

	// src := dir + "/" + filePath
	if boo := file.CheckPermission(path); boo {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", path)
	} 
    if err := file.IsNotExistMkDir(path); err != nil {
        return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", path, err)
    }
	f, err := file.Open(path + name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, fmt.Errorf("Fail to OpenFile :%v", err)
    }

    return f, nil
}
