package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func decodeImage(filename string) (image.Image, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

func main() {
	img, format, err := decodeImage("path/to/image.jpg")
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	fmt.Printf("Decoded image format: %s\n", format)
}
