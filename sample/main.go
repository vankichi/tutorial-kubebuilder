package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	path := "./configmap.yaml"
	f, err := os.OpenFile(path, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return
	}
	fmt.Println(f)
	// defer func() {
	// 	if f != nil {
	// 		if err != nil {
	// 			err = errors.Wrap(f.Close(), err.Error())
	// 			return
	// 		}
	// 		err = f.Close()
	// 	}
	// }()
	var cfg interface{}
	switch ext := filepath.Ext(path); ext {
	case ".yaml", ".yml":
		err = yaml.NewDecoder(f).Decode(cfg)
	default:
		// err = errors.ErrUnsupportedConfigFileType(ext)
	}
	fmt.Println(cfg)
	return
}
