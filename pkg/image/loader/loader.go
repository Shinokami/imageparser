package loader

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ImageParser(fileData []byte) imageList {
	localDir, _ := filepath.Abs("assets/images")
	images := imageList{}
	fileLines := strings.Split(string(fileData), "\n")
	for _, line := range fileLines {
		if line == "" {
			continue
		}
		rand.Seed(time.Now().UnixNano())
		fileExt := path.Ext(line)
		timestamp := time.Now().UnixNano()
		fileName := strings.Replace(path.Base(line), fileExt, "."+strconv.FormatInt(timestamp, 10)+fileExt, 1)
		images = append(images, image{
			url:       line,
			name:      fileName,
			localDir:  localDir,
			localPath: localDir + "/" + fileName})
	}

	return images
}

func LoadImages(images imageList) {
	// download
	// добавить создание папки
	for _, image := range images {
		res, err := http.Get(image.url)
		if err != nil {
			log.Printf("Error loading image", err)
			return
		}
		defer res.Body.Close()

		if _, err := os.Stat(image.localDir); os.IsNotExist(err) {
			err := os.Mkdir(image.localDir, 0755)
			if err != nil {
				log.Fatalf("Error creating folder for images: %s", err)
			}
		}

		file, err := os.Create(image.localPath)
		if err != nil {
			log.Fatalf("Error creating uploaded image file: %s", err)
		}
		// Получаем объект-читатель ответа на запрос на получение
		reader := bufio.NewReaderSize(res.Body, 32*1024)
		// Получить объект записи файла
		writer := bufio.NewWriter(file)
		io.Copy(writer, reader)
	}
}
