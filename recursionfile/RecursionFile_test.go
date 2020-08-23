package recursionfile

import "testing"

func TestRecursionFile(t *testing.T) {
	var files []string
	files, _ = RecursionFile("I:\\迅雷下载", files)
	for _, f := range files {
		t.Log(f)
	}
}
