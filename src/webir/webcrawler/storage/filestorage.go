package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	wch "webir/webcrawler/helper"

	log "github.com/sirupsen/logrus"
)

// FileStorage check localfile
type FileStorage struct {
	basePath       string
	totalFileCount int
}

// NewFileStorage create new file storage
func NewFileStorage(basePath string) *FileStorage {
	files, err := ioutil.ReadDir(basePath)
	var tc int
	if err == nil {
		tc = 0
	} else {
		tc = len(files)
	}
	return &FileStorage{basePath, tc}
}

func (fs *FileStorage) getPath(s string) string {
	return strings.Join([]string{fs.basePath, "/", s}, "")
}

// Write file
func (fs *FileStorage) Write(p string, c string, force bool) bool {
	if !force && fs.Exists(p) {
		log.Warn("File Exists not writing file")
		return false
	}
	fp := fs.getPath(p)
	dir := filepath.Dir(fp)
	// fn := filepath.Base(fp)
	log.Infof("creating directory: %v", dir)
	os.MkdirAll(dir, 0700)
	file, err := os.Create(fp)
	if err != nil {
		log.Warnf("Failed to create file: %v", fp)
		return false
	}
	n, err := file.WriteString(c)
	if err != nil {
		log.Warn("Failed to write file :<")
		return false
	}
	log.Debugf("Write %v bytes to file %v", n, p)
	log.Infof("Writted %v", fp)
	fs.totalFileCount++
	return true
}

// WriteURL file by URL
func (fs *FileStorage) WriteURL(u string, c string, force bool) bool {
	p := wch.URLtoFilePath(u)
	st := fs.Write(p, c, force)
	if !st {
		BLSet(u, InBlackList)
	}
	return st
}

func (fs *FileStorage) Read(p string) (string, error) {
	b, err := ioutil.ReadFile(fs.getPath(p))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadURL return file content
func (fs *FileStorage) ReadURL(u string) (string, error) {
	p := wch.URLtoFilePath(u)
	return fs.Read(p)
}

// Exists File Exists or not ?
func (fs *FileStorage) Exists(p string) bool {
	return wch.IsFileExists(fs.getPath(p))
}

// URLExists or not ?
func (fs *FileStorage) URLExists(u string) bool {
	p := wch.URLtoFilePath(u)
	return fs.Exists(p)
}

// URLFilePath ...
func (fs *FileStorage) URLFilePath(u string) string {
	p := wch.URLtoFilePath(u)
	return fs.getPath(p)
}

// GetTotalFileCount getter
func (fs *FileStorage) GetTotalFileCount() int {
	return fs.totalFileCount
}
