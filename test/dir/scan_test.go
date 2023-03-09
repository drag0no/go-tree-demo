package dir

import (
	"github.com/stretchr/testify/assert"
	"go-tree-demo/pkg/dir"
	"io/fs"
	"os"
	"testing"
)

type fileStatMock struct {
	Name     string
	IsDir    bool
	FileMode fs.FileMode
	FileInfo fs.FileInfo
}

type dirEntryMock struct {
	fs *fileStatMock
}

func (de dirEntryMock) Name() string               { return de.fs.Name }
func (de dirEntryMock) IsDir() bool                { return de.fs.IsDir }
func (de dirEntryMock) Type() fs.FileMode          { return de.fs.FileMode }
func (de dirEntryMock) Info() (fs.FileInfo, error) { return de.fs.FileInfo, nil }

var mockDirectory = dir.Directory{
	Name: ".",
	Directories: []*dir.Directory{
		{
			Name:        "Dir01",
			Directories: nil,
			Files: []*dir.File{
				{Name: "File01-01"},
			},
		},
		{
			Name:        "Dir02",
			Directories: nil,
			Files:       nil,
		},
	},
	Files: []*dir.File{
		{Name: "File00-01"},
		{Name: "File00-02"},
	},
}

func mockReadDir(name string) ([]os.DirEntry, error) {
	var dirEntry []os.DirEntry
	switch name {
	case ".":
		dirEntry = []os.DirEntry{
			dirEntryMock{&fileStatMock{
				Name:     "Dir01",
				IsDir:    true,
				FileMode: 0,
				FileInfo: nil,
			}},
			dirEntryMock{&fileStatMock{
				Name:     "Dir02",
				IsDir:    true,
				FileMode: 0,
				FileInfo: nil,
			}},
			dirEntryMock{&fileStatMock{
				Name:     "File00-01",
				IsDir:    false,
				FileMode: 0,
				FileInfo: nil,
			}},
			dirEntryMock{&fileStatMock{
				Name:     "File00-02",
				IsDir:    false,
				FileMode: 0,
				FileInfo: nil,
			}},
		}
	case "Dir01":
		dirEntry = []os.DirEntry{
			dirEntryMock{&fileStatMock{
				Name:     "File01-01",
				IsDir:    false,
				FileMode: 0,
				FileInfo: nil,
			}},
		}
	default:
		dirEntry = []os.DirEntry{}
	}
	return dirEntry, nil
}

func TestScanCustomHappy(t *testing.T) {
	expected := &mockDirectory
	actual := dir.ScanCustom(".", 5, mockReadDir)
	assert.Equal(t, expected, actual)
}

func TestScanNonExistentDir(t *testing.T) {
	var expected *dir.Directory
	actual := dir.Scan("random-duck", 5)
	assert.Equal(t, expected, actual)
}
