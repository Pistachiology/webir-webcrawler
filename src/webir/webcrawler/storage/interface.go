package storage

// Storage Interface
type Storage interface {
	Write(path string, c string, force bool) bool
	WriteURL(url string, c string, force bool) bool
	Read(path string) (string, error)
	ReadURL(url string) (string, error)
	Exists(p string) bool
	URLExists(u string) bool
	URLFilePath(u string) string
	GetTotalFileCount() int
}
