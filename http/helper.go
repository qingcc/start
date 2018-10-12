package app

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Time time.Time

func (c Time) String() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}

type Date time.Time

func (c Date) String() string {
	return time.Time(c).Format("2006-01-02")
}

//region Remark: 创建文件夹 Author; Qing
func DirectoryMkdir(path string) {
	if res, _ := DirectoryExists(path); res == false {
		err := os.MkdirAll(path, os.ModePerm)
		_, file, line, _ := runtime.Caller(0) // 获取错误文件和错误行
		println(file+":"+strconv.Itoa(line), err)
	}
}

//endregion

//region Remark: 判断文件夹是否存在 Author:Qing
func DirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//endregion

//region Remark: MD5加密 $ Author:Qing
func Strmd5(str string) string {
	w := md5.New()
	w.Write([]byte(str))
	return hex.EncodeToString(w.Sum(nil))
}

//endregion
