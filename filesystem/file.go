package filesystem

import (
	"errors"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
	"gopkg.in/go-mixed/framework.v1/facades"
	supportfile "gopkg.in/go-mixed/framework.v1/support/file"
	"gopkg.in/go-mixed/framework.v1/support/str"
)

type File struct {
	disk     string
	file     string
	filename string
}

func NewFile(file string) (*File, error) {
	if !supportfile.Exists(file) {
		return nil, errors.New("file doesn't exist")
	}

	disk := config.GetString("filesystems.default")

	return &File{disk: disk, file: file, filename: path.Base(file)}, nil
}

func NewFileFromRequest(fileHeader *multipart.FileHeader) (*File, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	tempFile, err := os.CreateTemp(os.TempDir(), "framework-")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, src)
	if err != nil {
		return nil, err
	}

	disk := config.GetString("filesystems.default")

	return &File{disk: disk, file: tempFile.Name(), filename: fileHeader.Filename}, nil
}

func (f *File) Disk(disk string) filesystem.File {
	f.disk = disk

	return f
}

func (f *File) File() string {
	return f.file
}

func (f *File) Store(path string) (string, error) {
	return facades.Storage.Disk(f.disk).PutFile(path, f)
}

func (f *File) StoreAs(path string, name string) (string, error) {
	return facades.Storage.Disk(f.disk).PutFileAs(path, f, name)
}

func (f *File) GetClientOriginalName() string {
	return f.filename
}

func (f *File) GetClientOriginalExtension() string {
	return supportfile.ClientOriginalExtension(f.filename)
}

func (f *File) HashName(path ...string) string {
	var realPath string
	if len(path) > 0 {
		realPath = strings.TrimRight(path[0], "/") + "/"
	}

	extension, _ := supportfile.Extension(f.file, true)
	if extension == "" {
		return realPath + str.Random(40)
	}

	return realPath + str.Random(40) + "." + extension
}

func (f *File) Extension() (string, error) {
	return supportfile.Extension(f.file)
}
