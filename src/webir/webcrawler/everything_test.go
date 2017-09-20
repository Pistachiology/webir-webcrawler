package webcrawler

import (
	"os"
	"reflect"
	"strings"
	"testing"
	wcf "webir/webcrawler/filter"
	wcfs "webir/webcrawler/storage"

	log "github.com/sirupsen/logrus"
)

// TestRemoveDuplicateFilter TEST
func TestRemoveDuplicateFilter(t *testing.T) {
	log.Info("TestRemoveDuplicateFilter")
	actual := wcf.RemoveDuplicate([]string{"aaa", "bbb", "aaa"})
	if !reflect.DeepEqual([]string{"aaa", "bbb"}, actual) {
		t.Errorf("FAIL")
	}
}

func TestExtensionFilter(t *testing.T) {
	extFunc := wcf.ExtensionFuncGenerator(nil)
	phpExtCheck := extFunc("http://www.facebook.com/a.php")
	falseExtCheck := extFunc("http://www.facebook.com/a.some")
	if falseExtCheck != false || phpExtCheck == false {
		t.Error()
	}
}

func TestCreateFileByFS(t *testing.T) {
	const storageDir = "test_output"
	const filename = "test_file"
	const fc = "Content"
	// write
	fs := wcfs.NewFileStorage(storageDir)
	written := fs.Write(filename, fc, false)
	if !written {
		t.Errorf("Couldn't write file :<")
	}
	// exists
	if !fs.Exists(filename) {
		t.Errorf("exists check bug")
	}

	// read
	p, err := fs.Read(filename)
	if err != nil {
		t.Errorf("Couldn't read file :<")
	}
	if strings.Compare(p, fc) != 0 {
		t.Error("File content not same :(")
	}
	os.RemoveAll(storageDir)
}

func TestCreateFileURLByFS(t *testing.T) {
	const storageDir = "test_output"
	const filename = "http://www.github.com/hello/abc.htm"
	const fc = "Content"
	// write
	fs := wcfs.NewFileStorage(storageDir)
	written := fs.WriteURL(filename, fc, true)
	if !written {
		t.Errorf("Couldn't write file :<")
	}
	// exists
	if !fs.URLExists("http://www.github.com/hello/abc.htm") {
		t.Errorf("exists check bug")
	}

	// read
	p, err := fs.ReadURL(filename)
	if err != nil {
		t.Errorf("Couldn't read file :<")
	}
	if strings.Compare(p, fc) != 0 {
		t.Error("File content not same :(")
	}
	os.RemoveAll(storageDir)
}
