package test11

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestNode(t *testing.T) {
	var arrS []string
	f, _ := GetAllFile(path.Join(`E:\钟情大美女，不亚于任何女明星，让人无比的渴望。最新PPV 031717_005-caribpr 無敵の無修正-丘咲エミリ`), arrS)
	for _, v := range f {
		t.Log(v)
	}
}

func GetAllFile(paths string, files []string) ([]string, error) {
	read, _ := ioutil.ReadDir(paths)
	for _, finfo := range read {
		if finfo.IsDir() {
			fullDir := path.Join(paths, finfo.Name())
			files = append(files, fullDir)
			files, _ = GetAllFile(fullDir, files)
		} else {
			fullDir := path.Join(paths, finfo.Name())
			files = append(files, fullDir)
		}
	}
	return files, nil
}
