package main

import (
	"io/ioutil"
	_ "io/ioutil"
	_ "log"
	"os"
	"path/filepath"
)

func main() {
	var (
		root  string
		files []string
		err   error
	)

	root = "./kartaca/"
	//filepath.Walk
	files, err = FilePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	//ioutil.ReadDir
	files, err = IOReadDir(root)
	if err != nil {
		panic(err)
	}
	//os.File.Readdir
	files, err = OSReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		//fmt.Println(file)
		os.Open(file)
	}
}

//Uygulama yolu içindeki dosyaları okur.
//func main() {
//	files, err := filepath.Glob("*")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(files) // contains a list of all files in the current directory
//}
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
