package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func CreateFile(name string, mode int64) (createFile *os.File, err error) {
	pathfile := string([]rune(name)[0:strings.LastIndex(name, "/")])
	fmt.Printf("name: ", name)
	fmt.Printf("pathfile: ", pathfile)
	err = os.MkdirAll(pathfile, 755)
	if err != nil {
		return nil, err
	}

	createFile, err = os.Create(name)
	fmt.Println("mode", mode, os.FileMode(mode))
	os.Chmod(name, os.FileMode(mode))
	return createFile, err
}

func DeCompress(tarFile, dest string) (err error) {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		fmt.Println(hdr)
		fmt.Println(hdr.Typeflag)
		filename := dest + hdr.Name
		fmt.Println("filename", filename, hdr.Mode)
		file, err := CreateFile(filename, hdr.Mode)
		if err != nil {
			fmt.Printf(err.Error())
		}
		io.Copy(file, tr)
	}
	return nil
}

func main() {
	tarfile := os.Args[1]
	dest := os.Args[2]
	DeCompress(tarfile, dest)
}
