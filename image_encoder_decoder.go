package main

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func DecodeImage(filename string) (image.Image, string, error) {
	file, err1 := os.Open(filename)
	if err1 != nil {
		return nil, "", err1
	}
	defer file.Close()

	img, format, err1 := image.Decode(file)
	if err1 != nil {
		return nil, "", err1
	}

	return img, format, nil
}

func EncodeImage(filename string, img image.Image, format string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(file, img, nil)
	case "png":
		return png.Encode(file, img)
	case "gif":
		return gif.Encode(file, img, nil)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

/*	*********************************Pour décoder l'image***********************
	img, format, err1 := decodeImage("path/to/image.jpg")
	if err1 != nil {
		fmt.Println("Error decoding image:", err1)
		return
	}
	fmt.Printf("Decoded image format: %s\n", format)

	*********************************Pour encoder l'image*****************
	err := encodeImage("path/to/output.jpg", img, format)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Image saved successfully") */
