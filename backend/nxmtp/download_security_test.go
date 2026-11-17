package nxmtp

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadDirectoryRejectsSymlink(t *testing.T) {
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(root, "outside")
	if err := os.Mkdir(outside, 0700); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "link")
	if err := os.Symlink(outside, link); err != nil {
		t.Fatal(err)
	}
	if err := makeDownloadDirectory(filepath.Join(link, "child")); err == nil {
		t.Fatal("followed symlink")
	}
	if _, err := os.Stat(filepath.Join(outside, "child")); !os.IsNotExist(err) {
		t.Fatal("created outside destination")
	}
	if err := makeDownloadDirectory(filepath.Join(root, "normal", "child")); err != nil {
		t.Fatal(err)
	}
}
