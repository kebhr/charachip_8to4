package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	os.Chdir(dir)

	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if strings.Contains(path, ".png") {
			fmt.Println(path)

			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			img, err := png.Decode(f)
			if err != nil {
				panic(err)
			}

			simg, ok := img.(SubImager)
			if ok {
				r := image.Rect(0, 0, img.Bounds().Dx()/2, img.Bounds().Dy())
				nimg := simg.SubImage(r)
				f.Close()

				out, err := os.Create(path)
				if err != nil {
					panic(err)
				}
				png.Encode(out, nimg)
				out.Close()
			} else {
				f.Close()
			}
		}
		return nil
	})
}
