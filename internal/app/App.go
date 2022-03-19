package app

import (
	"os"
	"path"

	"github.com/shinokami/imageparser/pkg/file/reader"
	"github.com/shinokami/imageparser/pkg/image/loader"
)

func Run() {
	localDir := "assets/txt"
	inputName := "input.txt"
	outputName := "output.txt"
	if len(os.Args) > 1 {
		inputName = os.Args[1]
		outputName = path.Dir(os.Args[1]) + "/" + outputName
	} else {
		inputName = localDir + "/" + inputName
		outputName = localDir + "/" + outputName
	}
	fileData := reader.ReadFile(inputName)
	images := loader.ImageParser(fileData)
	loader.LoadImages(images)
	reader.CreateFile(images.Paths(), outputName)
}
