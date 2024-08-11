package storage

import (
	"fmt"
	"os"
	"path"
)

func baseLocation() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, ".postman-alternative")
}
func Save(url string) error {
	os.MkdirAll(baseLocation(), os.ModePerm)

	file, err := os.Create(path.Join(baseLocation(), url+".toml"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString("Pruebaaaa")
	fmt.Println(err)
	file.Sync()
	return err
}

func loadDir(curr_path string) map[string]interface{} {
	root := make(map[string]interface{})
	filenames, err := os.ReadDir(curr_path)
	if err != nil {
		return nil
	}
	for _, fn := range filenames {
		if fn.IsDir() {
			newPath := path.Join(curr_path, fn.Name())
			root[fn.Name()] = loadDir(newPath)
		} else {
			root[fn.Name()] = fn.Name()
		}
	}
	return root
}
func LoadWorkspace() (map[string]interface{}, error) {
	return loadDir(baseLocation()), nil
}
