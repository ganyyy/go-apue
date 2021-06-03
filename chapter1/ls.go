// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("please input dir")
		return
	}

	var root = os.Args[1]

	// 必须指定一个root地址, 且打开的路径不能是 '/'
	// var dirFs = os.DirFS("")

	// // fs.ReadDirFS

	// fs.ValidPath()

	// var dir, err = fs.ReadDir(dirFs, root)
	// if err != nil {
	// 	log.Panicf("read %v error %v", root, err)
	// }

	// for _, d := range dir {
	// 	if d == nil {
	// 		continue
	// 	}
	// 	if d.IsDir() {
	// 		log.Printf("%v is dir", fmt.Sprintf("%s/%s", root, d.Name()))
	// 	} else {
	// 		log.Printf("%v is file", fmt.Sprintf("%s/%s", root, d.Name()))
	// 	}
	// }

	// fs.WalkDir(dirFs, os.Args[1], func(path string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		log.Printf("open %v error:%v", path, err)
	// 		return nil
	// 	}
	// 	if d.IsDir() {
	// 		log.Printf("%v is dir", path)
	// 	} else {
	// 		log.Printf("%v is file", path)
	// 	}
	// 	return nil
	// })
	var dir, err = os.ReadDir(root)
	if err != nil {
		log.Panicf("open %v error:%v", root, err)
	}
	if root == "/" {
		root = ""
	}
	for _, d := range dir {
		if d == nil {
			continue
		}
		fmt.Printf("%s/%s\n", root, d.Name())
	}
}
