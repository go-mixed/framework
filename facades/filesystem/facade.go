package filesystem

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
	"time"
)

func getFilesystem() filesystem.IStorage {
	return container.MustMakeAs("filesystem", filesystem.IStorage(nil))
}

func Disk(disk string) filesystem.IStorage {
	return getFilesystem().Disk(disk)
}

func AllDirectories(path string) ([]string, error) {
	return getFilesystem().AllDirectories(path)
}

func AllFiles(path string) ([]string, error) {
	return getFilesystem().AllFiles(path)
}

func Copy(oldFile, newFile string) error {
	return getFilesystem().Copy(oldFile, newFile)
}

func Delete(file ...string) error {
	return getFilesystem().Delete(file...)
}

func DeleteDirectory(directory string) error {
	return getFilesystem().DeleteDirectory(directory)
}

func Directories(path string) ([]string, error) {
	return getFilesystem().Directories(path)
}

// Download(path string)
func Exists(file string) bool {
	return getFilesystem().Exists(file)
}

func Files(path string) ([]string, error) {
	return getFilesystem().Files(path)
}

func Get(file string) (string, error) {
	return getFilesystem().Get(file)
}

func MakeDirectory(directory string) error {
	return getFilesystem().MakeDirectory(directory)
}

func Missing(file string) bool {
	return getFilesystem().Missing(file)
}

func Move(oldFile, newFile string) error {
	return getFilesystem().Move(oldFile, newFile)
}

func Path(file string) string {
	return getFilesystem().Path(file)
}

func Put(file, content string) error {
	return getFilesystem().Put(file, content)
}

func PutFile(path string, source filesystem.File) (string, error) {
	return getFilesystem().PutFile(path, source)
}

func PutFileAs(path string, source filesystem.File, name string) (string, error) {
	return getFilesystem().PutFileAs(path, source, name)
}

func Size(file string) (int64, error) {
	return getFilesystem().Size(file)
}

func TemporaryUrl(file string, time time.Time) (string, error) {
	return getFilesystem().TemporaryUrl(file, time)
}

func WithContext(ctx context.Context) filesystem.Driver {
	return getFilesystem().WithContext(ctx)
}

func Url(file string) string {
	return getFilesystem().Url(file)
}
