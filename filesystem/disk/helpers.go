package disk

import (
	"bytes"
	"io"
	"mime/multipart"
	"path"
	"strings"

	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
	"gopkg.in/go-mixed/framework.v1/support/file"
)

const MaxFileNum = 1000

func fileHeaderToString(fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}

func fullPathOfFile(filePath string, source filesystem.File, name string) (string, error) {
	extension := path.Ext(name)
	if extension == "" {
		var err error
		extension, err = file.Extension(source.File(), true)
		if err != nil {
			return "", err
		}

		return strings.TrimSuffix(filePath, "/") + "/" + strings.TrimSuffix(strings.TrimPrefix(path.Base(name), "/"), "/") + "." + extension, nil
	} else {
		extension = strings.TrimLeft(extension, ".")

		return strings.TrimSuffix(filePath, "/") + "/" + strings.TrimPrefix(path.Base(name), "/"), nil
	}
}

func validPath(path string) string {
	realPath := strings.TrimPrefix(path, "./")
	realPath = strings.TrimPrefix(realPath, "/")
	realPath = strings.TrimPrefix(realPath, ".")
	if realPath != "" && !strings.HasSuffix(realPath, "/") {
		realPath += "/"
	}

	return realPath
}
