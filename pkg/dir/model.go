package dir

type File struct {
	Name string
}

type Directory struct {
	Name        string
	Directories []*Directory
	Files       []*File
}
