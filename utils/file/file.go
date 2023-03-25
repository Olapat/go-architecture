package fileUtils

import (
	"fmt"
	"mime/multipart"
	"os"
	"regexp"
)

type FileUtil struct {
	FileCtx *multipart.FileHeader
}

func (fileUtil FileUtil) CheckType(allowedTypes []string) bool {
	fType := fileUtil.FileCtx.Header.Get("Content-Type")
	checkedType := false
	for _, at := range allowedTypes {
		if at == fType {
			checkedType = true
			break
		}
	}

	return checkedType
}

func GetPathFile(path string) string {
	host := os.Getenv("HOST")
	var re = regexp.MustCompile(`(?m)^http`)
	var rf = re.FindAllString(path, -1)
	if len(rf) > 0 {
		return path
	}
	var avp = fmt.Sprintf("%s/file%s", host, path)
	return avp
}
