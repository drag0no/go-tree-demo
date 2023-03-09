package dir

import (
	"strings"
)

type offsetFn func(int) string
type formatFn func(*strings.Builder, string, string)

func AsTreeStr(dir *Directory) string {
	return AsTreeStrCustom(dir, defaultOffsetFn, defaultFormatFn)
}

func AsTreeStrCustom(dir *Directory, offsetFunc offsetFn, formatFunc formatFn) string {
	builder := strings.Builder{}
	dirToTreeStr(&builder, dir, offsetFunc, formatFunc, 0)
	return builder.String()
}

func dirToTreeStr(
	builder *strings.Builder,
	dir *Directory,
	offsetFunc offsetFn,
	formatFunc formatFn,
	level int,
) {
	if dir == nil {
		return
	}

	offset := ""
	if level != 0 {
		offset = offsetFunc(level)
	}

	formatFunc(builder, dir.Name, offset)

	if dir.Directories != nil {
		for _, entry := range dir.Directories {
			dirToTreeStr(builder, entry, defaultOffsetFn, defaultFormatFn, level+1)
		}
	}
	if dir.Files != nil {
		for _, entry := range dir.Files {
			fileToTreeStr(builder, entry, defaultOffsetFn, defaultFormatFn, level+1)
		}
	}
}

func fileToTreeStr(
	builder *strings.Builder,
	file *File,
	offsetFunc offsetFn,
	formatFunc formatFn, level int,
) {
	if file == nil {
		return
	}
	offset := offsetFunc(level)
	formatFunc(builder, file.Name, offset)
}

func defaultOffsetFn(level int) string {
	return strings.Repeat("| ", level-1) + "├─"
}

func defaultFormatFn(
	builder *strings.Builder,
	name string,
	offset string,
) {
	builder.WriteString(offset + name + "\n")
}
