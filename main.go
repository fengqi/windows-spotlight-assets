package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	destination := os.Getenv("USERPROFILE")
	destination += "\\Pictures\\自带壁纸"
	log.Println("Destination: " + destination)
	if !exists(destination) {
		_ = os.Mkdir(destination, 0664)
	}

	path := os.Getenv("LocalAppData")
	path += "\\Packages\\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\\LocalState\\Assets"
	log.Println("Assets Path: " + path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("ReadDir err: " + err.Error())
		return
	}

	for _, f := range files {
		dst := destination + "\\" + f.Name() + ".jpg"
		if exists(dst) {
			continue
		}

		src := path + "\\" + f.Name()
		input, err := ioutil.ReadFile(src)
		if err != nil {
			log.Println("ReadFile err: " + err.Error())
			continue
		}

		img, _, err := image.DecodeConfig(bytes.NewReader(input))
		if err != nil || img.Height > img.Width {
			continue
		}

		err = ioutil.WriteFile(dst, input, 0664)
		if err != nil {
			log.Println("CopyFile err: " + err.Error())
		} else {
			log.Println("CopyFile: " + f.Name() + "to " + dst)
		}
	}
}

func exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil || os.IsExist(err)
}
