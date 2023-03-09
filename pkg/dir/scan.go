package dir

import (
	"log"
	"os"
	"path/filepath"
)

type ReadDirFn func(name string) ([]os.DirEntry, error)

func Scan(path string, maxDepth int) *Directory {
	return ScanCustom(path, maxDepth, os.ReadDir)
}

func ScanCustom(path string, maxDepth int, readDirFunc ReadDirFn) *Directory {
	return scan(path, 0, maxDepth, readDirFunc)
}

func scan(path string, curDepth int, maxDepth int, readDirFunc ReadDirFn) *Directory {
	curDepth++
	var directoryList []*Directory
	var fileList []*File

	readDirResult, err := readDirFunc(path)
	if err != nil {
		log.Println(err)
		return nil
	}

	if curDepth <= maxDepth {
		for _, entry := range readDirResult {
			if entry.IsDir() {
				dir := scan(filepath.Join(path, entry.Name()), curDepth, maxDepth, readDirFunc)
				if dir != nil {
					directoryList = append(directoryList, dir)
				}
			} else {
				fileList = append(fileList, &File{entry.Name()})
			}
		}
	}

	return &Directory{filepath.Base(path), directoryList, fileList}
}
