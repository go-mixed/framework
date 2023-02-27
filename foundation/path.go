package foundation

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func getCurrentAbsPath() string {
	dir := getCurrentAbsPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbsPathByCaller()
	}

	return dir
}

func getCurrentAbsPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))

	return res
}

func getCurrentAbsPathByCaller() string {
	var abPath string
	for i := 0; i < 15; i++ {
		_, filename, _, ok := runtime.Caller(i)
		if ok && strings.HasSuffix(filename, "main.go") {
			abPath = path.Dir(filename)
			break
		}
	}

	return abPath
}
