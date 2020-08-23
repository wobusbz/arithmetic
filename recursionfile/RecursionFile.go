package recursionfile

import "io/ioutil"

func RecursionFile(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "\\" + fi.Name()
			files = append(files, fulldir)
			files, _ = RecursionFile(fulldir, files)
		} else {
			fulldir := path + "\\" + fi.Name()
			files = append(files, fulldir)
		}
	}
	return files, nil
}
