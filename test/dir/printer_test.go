package dir

import (
	"github.com/drag0no/go-tree-demo/pkg/dir"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsTreeStrHappy(t *testing.T) {
	directory := dir.Directory{
		Name: ".",
		Directories: []*dir.Directory{
			{
				Name: "Dir01",
				Directories: []*dir.Directory{
					{
						Name:        "Dir03",
						Directories: nil,
						Files:       nil,
					},
				},
				Files: []*dir.File{
					{Name: "File01-01"},
					{Name: "File01-02"},
				},
			},
			{
				Name:        "Dir02",
				Directories: nil,
				Files: []*dir.File{
					{Name: "File02-01"},
				},
			},
		},
		Files: []*dir.File{
			{Name: "File00-01"},
			{Name: "File00-02"},
			{Name: "File00-03"},
			nil,
		},
	}
	expected := `.
├─Dir01
| ├─Dir03
| ├─File01-01
| ├─File01-02
├─Dir02
| ├─File02-01
├─File00-01
├─File00-02
├─File00-03
`
	actual := dir.AsTreeStr(&directory)
	assert.Equal(t, expected, actual)
}

func TestAsTreeStrNilDir(t *testing.T) {
	var directory *dir.Directory = nil
	expected := ""
	actual := dir.AsTreeStr(directory)
	assert.Equal(t, expected, actual)
}
