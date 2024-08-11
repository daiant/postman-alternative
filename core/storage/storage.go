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
